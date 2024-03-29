// Copyright 2019 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package minifiers

import (
	"github.com/spf13/cast"
	"github.com/strawberryssg/strawberry-v0/common/maps"
	"github.com/strawberryssg/strawberry-v0/config"
	"github.com/strawberryssg/strawberry-v0/docshelper"
	"github.com/strawberryssg/strawberry-v0/parser"

	"github.com/mitchellh/mapstructure"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

var defaultTdewolffConfig = tdewolffConfig{
	HTML: html.Minifier{
		KeepDocumentTags:        true,
		KeepConditionalComments: true,
		KeepEndTags:             true,
		KeepDefaultAttrVals:     true,
		KeepWhitespace:          false,
	},
	CSS: css.Minifier{
		Precision: 0,
		KeepCSS2:  true,
	},
	JS:   js.Minifier{},
	JSON: json.Minifier{},
	SVG: svg.Minifier{
		KeepComments: false,
		Precision:    0,
	},
	XML: xml.Minifier{
		KeepWhitespace: false,
	},
}

type tdewolffConfig struct {
	HTML html.Minifier
	CSS  css.Minifier
	JS   js.Minifier
	JSON json.Minifier
	SVG  svg.Minifier
	XML  xml.Minifier
}

type minifyConfig struct {
	// Whether to minify the published output (the HTML written to /public).
	MinifyOutput bool

	DisableHTML bool
	DisableCSS  bool
	DisableJS   bool
	DisableJSON bool
	DisableSVG  bool
	DisableXML  bool

	Tdewolff tdewolffConfig
}

var defaultConfig = minifyConfig{
	Tdewolff: defaultTdewolffConfig,
}

func decodeConfig(cfg config.Provider) (conf minifyConfig, err error) {
	conf = defaultConfig

	// May be set by CLI.
	conf.MinifyOutput = cfg.GetBool("minifyOutput")

	v := cfg.Get("minify")
	if v == nil {
		return
	}

	// Legacy.
	if b, ok := v.(bool); ok {
		conf.MinifyOutput = b
		return
	}

	m := maps.ToStringMap(v)

	// Handle upstream renames.
	if td, found := m["tdewolff"]; found {
		tdm := maps.ToStringMap(td)
		for _, key := range []string{"css", "svg"} {
			if v, found := tdm[key]; found {
				vm := maps.ToStringMap(v)
				if vv, found := vm["decimal"]; found {
					vvi := cast.ToInt(vv)
					if vvi > 0 {
						vm["precision"] = vvi
					}
				}
			}
		}
	}

	err = mapstructure.WeakDecode(m, &conf)

	if err != nil {
		return
	}

	return
}

func init() {
	docsProvider := func() docshelper.DocProvider {
		return docshelper.DocProvider{"config": map[string]any{"minify": parser.LowerCaseCamelJSONMarshaller{Value: defaultConfig}}}
	}
	docshelper.AddDocProviderFunc(docsProvider)
}
