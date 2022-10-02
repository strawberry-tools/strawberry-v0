package media

import (
	"github.com/strawberryssg/strawberry-v0/docshelper"
)

// This is is just some helpers used to create some JSON used in the Hugo docs.
func init() {
	docsProvider := func() docshelper.DocProvider {
		return docshelper.DocProvider{"media": map[string]any{"types": DefaultTypes}}
	}
	docshelper.AddDocProviderFunc(docsProvider)
}
