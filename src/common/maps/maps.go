// Copyright 2018 The Hugo Authors. All rights reserved.
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

package maps

import (
	"fmt"
	"strings"

	"github.com/strawberryssg/strawberry-v0/common/types"

	"github.com/gobwas/glob"
	"github.com/spf13/cast"
)

// ToStringMapE converts in to map[string]interface{}.
func ToStringMapE(in any) (map[string]any, error) {
	switch vv := in.(type) {
	case Params:
		return vv, nil
	case map[string]string:
		var m = map[string]any{}
		for k, v := range vv {
			m[k] = v
		}
		return m, nil

	default:
		return cast.ToStringMapE(in)
	}
}

// ToParamsAndPrepare converts in to Params and prepares it for use.
// If in is nil, an empty map is returned.
// See PrepareParams.
func ToParamsAndPrepare(in any) (Params, bool) {
	if types.IsNil(in) {
		return Params{}, true
	}
	m, err := ToStringMapE(in)
	if err != nil {
		return nil, false
	}
	PrepareParams(m)
	return m, true
}

// MustToParamsAndPrepare calls ToParamsAndPrepare and panics if it fails.
func MustToParamsAndPrepare(in any) Params {
	if p, ok := ToParamsAndPrepare(in); ok {
		return p
	} else {
		panic(fmt.Sprintf("cannot convert %T to maps.Params", in))
	}
}

// ToStringMap converts in to map[string]interface{}.
func ToStringMap(in any) map[string]any {
	m, _ := ToStringMapE(in)
	return m
}

// ToStringMapStringE converts in to map[string]string.
func ToStringMapStringE(in any) (map[string]string, error) {
	m, err := ToStringMapE(in)
	if err != nil {
		return nil, err
	}
	return cast.ToStringMapStringE(m)
}

// ToStringMapString converts in to map[string]string.
func ToStringMapString(in any) map[string]string {
	m, _ := ToStringMapStringE(in)
	return m
}

// ToStringMapBool converts in to bool.
func ToStringMapBool(in any) map[string]bool {
	m, _ := ToStringMapE(in)
	return cast.ToStringMapBool(m)
}

// ToSliceStringMap converts in to []map[string]interface{}.
func ToSliceStringMap(in any) ([]map[string]any, error) {
	switch v := in.(type) {
	case []map[string]any:
		return v, nil
	case []any:
		var s []map[string]any
		for _, entry := range v {
			if vv, ok := entry.(map[string]any); ok {
				s = append(s, vv)
			}
		}
		return s, nil
	default:
		return nil, fmt.Errorf("unable to cast %#v of type %T to []map[string]interface{}", in, in)
	}
}

// LookupEqualFold finds key in m with case insensitive equality checks.
func LookupEqualFold[T any | string](m map[string]T, key string) (T, bool) {
	if v, found := m[key]; found {
		return v, true
	}
	for k, v := range m {
		if strings.EqualFold(k, key) {
			return v, true
		}
	}
	var s T
	return s, false
}

type keyRename struct {
	pattern glob.Glob
	newKey  string
}

// KeyRenamer supports renaming of keys in a map.
type KeyRenamer struct {
	renames []keyRename
}

// NewKeyRenamer creates a new KeyRenamer given a list of pattern and new key
// value pairs.
func NewKeyRenamer(patternKeys ...string) (KeyRenamer, error) {
	var renames []keyRename
	for i := 0; i < len(patternKeys); i += 2 {
		g, err := glob.Compile(strings.ToLower(patternKeys[i]), '/')
		if err != nil {
			return KeyRenamer{}, err
		}
		renames = append(renames, keyRename{pattern: g, newKey: patternKeys[i+1]})
	}

	return KeyRenamer{renames: renames}, nil
}

func (r KeyRenamer) getNewKey(keyPath string) string {
	for _, matcher := range r.renames {
		if matcher.pattern.Match(keyPath) {
			return matcher.newKey
		}
	}

	return ""
}

// Rename renames the keys in the given map according
// to the patterns in the current KeyRenamer.
func (r KeyRenamer) Rename(m map[string]any) {
	r.renamePath("", m)
}

func (KeyRenamer) keyPath(k1, k2 string) string {
	k1, k2 = strings.ToLower(k1), strings.ToLower(k2)
	if k1 == "" {
		return k2
	}
	return k1 + "/" + k2
}

func (r KeyRenamer) renamePath(parentKeyPath string, m map[string]any) {
	for key, val := range m {
		keyPath := r.keyPath(parentKeyPath, key)
		switch val.(type) {
		case map[any]any:
			val = cast.ToStringMap(val)
			r.renamePath(keyPath, val.(map[string]any))
		case map[string]any:
			r.renamePath(keyPath, val.(map[string]any))
		}

		newKey := r.getNewKey(keyPath)

		if newKey != "" {
			delete(m, key)
			m[newKey] = val
		}
	}
}
