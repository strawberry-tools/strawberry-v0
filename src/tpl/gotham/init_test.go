// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package gotham

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/strawberryssg/strawberry-v0/deps"
	"github.com/strawberryssg/strawberry-v0/htesting/hqt"
	"github.com/strawberryssg/strawberry-v0/tpl/internal"
)

func TestInit(t *testing.T) {
	c := qt.New(t)
	var found bool
	var ns *internal.TemplateFuncsNamespace

	for _, nsf := range internal.TemplateFuncsNamespaceRegistry {
		ns = nsf(&deps.Deps{})
		if ns.Name == name {
			found = true
			break
		}
	}

	c.Assert(found, qt.Equals, true)

	ctx, err := ns.Context()
	c.Assert(err, qt.IsNil)
	c.Assert(ctx, hqt.IsSameType, &Namespace{})
}
