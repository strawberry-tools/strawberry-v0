// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package gotham

import (
	"html/template"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestQRCoder(t *testing.T) {

	t.Parallel()
	c := qt.New(t)
	ns := New()

	for i, test := range []struct {
		in     interface{}
		expect interface{}
	}{
		{"https://google.com", template.HTML(`data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABZ0lEQVR42uyYwc3sIAyEHXHgSAkphdZSGqVQAkcOiHkaQ1bJ+1dbAGYOK23ynax4xrZsbW1t/RJUzVXxQJGQz/HkWgvo/HHNVV99EQKZT6I54GCZmrBOYKlOlipZBVzV70XEOCC+igQAeVFg9IXwgyihsA5fG2d14OOT9IdQfhjp2sCUApit8U2rA/1Ad0DzUIfQvIgp4loLONC1MTwq/4asjRGTGAP43mHE4vAH5Pc0aAQAumsOVXwNQOH7iGQPYHAOn9T/J/KZ4tNIlwDYGU0cE6CGEgrTID1j0Qhw9IM+SXdgLN6FehiIDaDLwb7gvMg5ifPiu1BGAP1gxpbEzgnAnzFpDWBOQToWMxblTDFFa8DzSDKPA/8ZqQ3gPpIwF4sEaKEk2gPurXmcSEbjpPd2sA7wWYIYB+/boDFgjMV6JnnHohFgHknmlsQdiT4p1oDbJ4E6xqR5JbnWAra2tizqXwAAAP//EnuWL6vnRrMAAAAASUVORK5CYII=`)},
		{"https://gothamhq.com", template.HTML(`data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQAAAAEAAQMAAABmvDolAAAABlBMVEX///8AAABVwtN+AAABYklEQVR42uyYMa4DIQxEjSgoOcIehav9o3GUPQIlBWK+bJMIpGz6YKaIouVVFjM2pqOjo6NvgqhRqAElotClX/72Ajr/+ObBAMVC181fkjnAAc03/hsGcAHZLhBQ+dw44AHUCODeFFBfUADXQYCPxtkdGDnpa6ihxBLvhyDdHZgUAJRvHXJnoLvuNCdDjXp8ZUp5M4DIdTEGcVsUZ6y+MAI4oHuekyhUYoADIk2+sAJ0nZJGTI4xabowRgBww/B8LGXizkkppylItwD4Rmg66H2Q65BX41gAOpFrHu09DEql1leSBcCNOUkLJTG59gsjQHeSk5wP/ErSGSkv48EOwJCvPASVyE3xwpoPJoDXkmSc6yvJIPB6Jcm2CCWOQk2+MAJMr2aKRZ2zFGonQBqj9ANZDmYim4DXOpHch/xhN7g98PYFtC3ifjLO1sCyJJFHEreM9LBF+Vng6OjIov4DAAD//wwh/kPBEpK+AAAAAElFTkSuQmCC`)},
		// error
		{"", false},
	} {
		errMsg := qt.Commentf("[%d] %v", i, test.in)

		result, err := ns.QRCoder(test.in)

		if b, ok := test.expect.(bool); ok && !b {
			c.Assert(err, qt.Not(qt.IsNil), errMsg)
			continue
		}

		c.Assert(err, qt.IsNil, errMsg)
		c.Assert(result, qt.Equals, test.expect, errMsg)
	}
}
