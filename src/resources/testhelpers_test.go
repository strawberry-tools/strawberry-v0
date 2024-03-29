package resources

import (
	"image"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/strawberryssg/strawberry-v0/cache/filecache"
	"github.com/strawberryssg/strawberry-v0/config"
	"github.com/strawberryssg/strawberry-v0/helpers"
	"github.com/strawberryssg/strawberry-v0/hugofs"
	"github.com/strawberryssg/strawberry-v0/langs"
	"github.com/strawberryssg/strawberry-v0/media"
	"github.com/strawberryssg/strawberry-v0/modules"
	"github.com/strawberryssg/strawberry-v0/output"
	"github.com/strawberryssg/strawberry-v0/resources/images"
	"github.com/strawberryssg/strawberry-v0/resources/page"
	"github.com/strawberryssg/strawberry-v0/resources/resource"

	"github.com/spf13/afero"

	qt "github.com/frankban/quicktest"
)

type specDescriptor struct {
	baseURL string
	c       *qt.C
	fs      afero.Fs
}

func createTestCfg() config.Provider {
	cfg := config.New()
	cfg.Set("resourceDir", "resources")
	cfg.Set("contentDir", "content")
	cfg.Set("dataDir", "data")
	cfg.Set("i18nDir", "i18n")
	cfg.Set("layoutDir", "layouts")
	cfg.Set("assetDir", "assets")
	cfg.Set("archetypeDir", "archetypes")
	cfg.Set("publishDir", "public")

	langs.LoadLanguageSettings(cfg, nil)
	mod, err := modules.CreateProjectModule(cfg)
	if err != nil {
		panic(err)
	}
	cfg.Set("allModules", modules.Modules{mod})

	return cfg
}

func newTestResourceSpec(desc specDescriptor) *Spec {
	baseURL := desc.baseURL
	if baseURL == "" {
		baseURL = "https://example.com/"
	}

	afs := desc.fs
	if afs == nil {
		afs = afero.NewMemMapFs()
	}

	afs = hugofs.NewBaseFileDecorator(afs)

	c := desc.c

	cfg := createTestCfg()
	cfg.Set("baseURL", baseURL)

	imagingCfg := map[string]any{
		"resampleFilter": "linear",
		"quality":        68,
		"anchor":         "left",
	}

	cfg.Set("imaging", imagingCfg)

	fs := hugofs.NewFrom(afs, cfg)
	fs.PublishDir = hugofs.NewCreateCountingFs(fs.PublishDir)

	s, err := helpers.NewPathSpec(fs, cfg, nil)
	c.Assert(err, qt.IsNil)

	filecaches, err := filecache.NewCaches(s)
	c.Assert(err, qt.IsNil)

	spec, err := NewSpec(s, filecaches, nil, nil, nil, nil, output.DefaultFormats, media.DefaultTypes)
	c.Assert(err, qt.IsNil)
	return spec
}

func newTargetPaths(link string) func() page.TargetPaths {
	return func() page.TargetPaths {
		return page.TargetPaths{
			SubResourceBaseTarget: filepath.FromSlash(link),
			SubResourceBaseLink:   link,
		}
	}
}

func newTestResourceOsFs(c *qt.C) (*Spec, string) {
	cfg := createTestCfg()
	cfg.Set("baseURL", "https://example.com")

	workDir, err := ioutil.TempDir("", "hugores")
	c.Assert(err, qt.IsNil)
	c.Assert(workDir, qt.Not(qt.Equals), "")

	if runtime.GOOS == "darwin" && !strings.HasPrefix(workDir, "/private") {
		// To get the entry folder in line with the rest. This its a little bit
		// mysterious, but so be it.
		workDir = "/private" + workDir
	}

	cfg.Set("workingDir", workDir)

	fs := hugofs.NewFrom(hugofs.NewBaseFileDecorator(hugofs.Os), cfg)

	s, err := helpers.NewPathSpec(fs, cfg, nil)
	c.Assert(err, qt.IsNil)

	filecaches, err := filecache.NewCaches(s)
	c.Assert(err, qt.IsNil)

	spec, err := NewSpec(s, filecaches, nil, nil, nil, nil, output.DefaultFormats, media.DefaultTypes)
	c.Assert(err, qt.IsNil)

	return spec, workDir
}

func fetchSunset(c *qt.C) images.ImageResource {
	return fetchImage(c, "sunset.jpg")
}

func fetchImage(c *qt.C, name string) images.ImageResource {
	spec := newTestResourceSpec(specDescriptor{c: c})
	return fetchImageForSpec(spec, c, name)
}

func fetchImageForSpec(spec *Spec, c *qt.C, name string) images.ImageResource {
	r := fetchResourceForSpec(spec, c, name)

	img := r.(images.ImageResource)

	c.Assert(img, qt.Not(qt.IsNil))
	c.Assert(img.(specProvider).getSpec(), qt.Not(qt.IsNil))

	return img
}

func fetchResourceForSpec(spec *Spec, c *qt.C, name string, targetPathAddends ...string) resource.ContentResource {
	src, err := os.Open(filepath.FromSlash("testdata/" + name))
	c.Assert(err, qt.IsNil)
	workDir := spec.WorkingDir
	if len(targetPathAddends) > 0 {
		addends := strings.Join(targetPathAddends, "_")
		name = addends + "_" + name
	}
	targetFilename := filepath.Join(workDir, name)
	out, err := helpers.OpenFileForWriting(spec.Fs.Source, targetFilename)
	c.Assert(err, qt.IsNil)
	_, err = io.Copy(out, src)
	out.Close()
	src.Close()
	c.Assert(err, qt.IsNil)

	factory := newTargetPaths("/a")

	r, err := spec.New(ResourceSourceDescriptor{Fs: spec.Fs.Source, TargetPaths: factory, LazyPublish: true, RelTargetFilename: name, SourceFilename: targetFilename})
	c.Assert(err, qt.IsNil)
	c.Assert(r, qt.Not(qt.IsNil))

	return r.(resource.ContentResource)
}

func assertImageFile(c *qt.C, fs afero.Fs, filename string, width, height int) {
	filename = filepath.Clean(filename)
	f, err := fs.Open(filename)
	c.Assert(err, qt.IsNil)
	defer f.Close()

	config, _, err := image.DecodeConfig(f)
	c.Assert(err, qt.IsNil)

	c.Assert(config.Width, qt.Equals, width)
	c.Assert(config.Height, qt.Equals, height)
}

func assertFileCache(c *qt.C, fs afero.Fs, filename string, width, height int) {
	assertImageFile(c, fs, filepath.Clean(filename), width, height)
}

func writeSource(t testing.TB, fs *hugofs.Fs, filename, content string) {
	writeToFs(t, fs.Source, filename, content)
}

func writeToFs(t testing.TB, fs afero.Fs, filename, content string) {
	if err := afero.WriteFile(fs, filepath.FromSlash(filename), []byte(content), 0755); err != nil {
		t.Fatalf("Failed to write file: %s", err)
	}
}
