// Copyright 2020 The Hugo Authors. All rights reserved.
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

package openapi3

import (
	"fmt"
	"io/ioutil"

	gyaml "github.com/ghodss/yaml"

	"errors"

	kopenapi3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/strawberryssg/strawberry-v0/cache/namedmemcache"
	"github.com/strawberryssg/strawberry-v0/deps"
	"github.com/strawberryssg/strawberry-v0/parser/metadecoders"
	"github.com/strawberryssg/strawberry-v0/resources/resource"
)

// New returns a new instance of the openapi3-namespaced template functions.
func New(deps *deps.Deps) *Namespace {
	// TODO(bep) consolidate when merging that "other branch" -- but be aware of the keys.
	cache := namedmemcache.New()
	deps.BuildStartListeners.Add(
		func() {
			cache.Clear()
		})

	return &Namespace{
		cache: cache,
		deps:  deps,
	}
}

// Namespace provides template functions for the "openapi3".
type Namespace struct {
	cache *namedmemcache.Cache
	deps  *deps.Deps
}

func (ns *Namespace) Unmarshal(r resource.UnmarshableResource) (*kopenapi3.T, error) {
	key := r.Key()
	if key == "" {
		return nil, errors.New("no Key set in Resource")
	}

	v, err := ns.cache.GetOrCreate(key, func() (any, error) {
		f := metadecoders.FormatFromMediaType(r.MediaType())
		if f == "" {
			return nil, fmt.Errorf("MIME %q not supported", r.MediaType())
		}

		reader, err := r.ReadSeekCloser()
		if err != nil {
			return nil, err
		}
		defer reader.Close()

		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		s := &kopenapi3.T{}
		switch f {
		case metadecoders.YAML:
			err = gyaml.Unmarshal(b, s)
		default:
			err = metadecoders.Default.UnmarshalTo(b, f, s)
		}
		if err != nil {
			return nil, err
		}

		err = kopenapi3.NewLoader().ResolveRefsIn(s, nil)

		return s, err
	})
	if err != nil {
		return nil, err
	}

	return v.(*kopenapi3.T), nil
}
