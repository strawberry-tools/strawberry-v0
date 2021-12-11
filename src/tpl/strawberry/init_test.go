// Copyright 2021 Strawberry Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package strawberry

import (
	"testing"

	"github.com/strawberryssg/strawberry-v0/config"
	"github.com/strawberryssg/strawberry-v0/deps"
	"github.com/strawberryssg/strawberry-v0/htesting/hqt"
	"github.com/strawberryssg/strawberry-v0/resources/page"
	"github.com/strawberryssg/strawberry-v0/tpl/internal"

	qt "github.com/frankban/quicktest"
)

func TestInit(t *testing.T) {
	c := qt.New(t)
	var found bool
	var ns *internal.TemplateFuncsNamespace
	v := config.New()
	v.Set("contentDir", "content")
	s := page.NewDummyHugoSite(v)

	for _, nsf := range internal.TemplateFuncsNamespaceRegistry {
		ns = nsf(&deps.Deps{Site: s})
		if ns.Name == name {
			found = true
			break
		}
	}

	c.Assert(found, qt.Equals, true)
	ctx, err := ns.Context()
	c.Assert(err, qt.IsNil)
	c.Assert(ctx, hqt.IsSameType, s.Strawberry())
}
