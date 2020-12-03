// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package gotham

import (
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/gothamhq/gotham/deps"
	"github.com/gothamhq/gotham/htesting/hqt"
	"github.com/gothamhq/gotham/tpl/internal"
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
	c.Assert(ns.Context(), hqt.IsSameType, &Namespace{})
}
