// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package gotham

import (
	"github.com/gothamhq/gotham/deps"
	"github.com/gothamhq/gotham/tpl/internal"
)

const name = "gotham"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		ctx := New()

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) interface{} { return ctx },
		}

		ns.AddMethodMapping(ctx.QRCoder,
			[]string{"qrcoder"},
			[][2]string{
				{`{{ qrcoder "https://google.com" }}`, `data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABZ0lEQVR42uyYwc3sIAyEHXHgSAkphdZSGqVQAkcOiHkaQ1bJ+1dbAGYOK23ynax4xrZsbW1t/RJUzVXxQJGQz/HkWgvo/HHNVV99EQKZT6I54GCZmrBOYKlOlipZBVzV70XEOCC+igQAeVFg9IXwgyihsA5fG2d14OOT9IdQfhjp2sCUApit8U2rA/1Ad0DzUIfQvIgp4loLONC1MTwq/4asjRGTGAP43mHE4vAH5Pc0aAQAumsOVXwNQOH7iGQPYHAOn9T/J/KZ4tNIlwDYGU0cE6CGEgrTID1j0Qhw9IM+SXdgLN6FehiIDaDLwb7gvMg5ifPiu1BGAP1gxpbEzgnAnzFpDWBOQToWMxblTDFFa8DzSDKPA/8ZqQ3gPpIwF4sEaKEk2gPurXmcSEbjpPd2sA7wWYIYB+/boDFgjMV6JnnHohFgHknmlsQdiT4p1oDbJ4E6xqR5JbnWAra2tizqXwAAAP//EnuWL6vnRrMAAAAASUVORK5CYII=`},
			},
		)

		return ns

	}

	internal.AddTemplateFuncsNamespace(f)
}
