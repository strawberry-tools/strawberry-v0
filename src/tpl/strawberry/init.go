// Copyright 2021 Strawberry Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Package strawberry mirrors the hugo package. It provides Go template
// functions for Strawberry/Hugo metadata.
package strawberry

import (
	"github.com/strawberryssg/strawberry-v0/deps"
	"github.com/strawberryssg/strawberry-v0/tpl/internal"
)

const name = "strawberry"

func init() {
	f := func(d *deps.Deps) *internal.TemplateFuncsNamespace {
		h := d.Site.Strawberry()

		ns := &internal.TemplateFuncsNamespace{
			Name:    name,
			Context: func(args ...interface{}) (interface{}, error) { return h, nil },
		}

		// We just add the Strawberry struct as the namespace here. No method mappings.

		return ns
	}

	internal.AddTemplateFuncsNamespace(f)
}
