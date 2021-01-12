// Copyright 2017-present The Hugo Authors. All rights reserved.
//
// Portions Copyright The Go Authors.

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

package tplimpl

import (
	"reflect"
	"strings"

	"github.com/gothamhq/gotham/tpl"

	"github.com/gothamhq/gotham/common/maps"

	template "github.com/gothamhq/gotham/tpl/internal/go_templates/htmltemplate"
	texttemplate "github.com/gothamhq/gotham/tpl/internal/go_templates/texttemplate"

	"github.com/gothamhq/gotham/deps"

	"github.com/gothamhq/gotham/tpl/internal"

	// Init the namespaces
	_ "github.com/gothamhq/gotham/tpl/cast"
	_ "github.com/gothamhq/gotham/tpl/collections"
	_ "github.com/gothamhq/gotham/tpl/compare"
	_ "github.com/gothamhq/gotham/tpl/crypto"
	_ "github.com/gothamhq/gotham/tpl/data"
	_ "github.com/gothamhq/gotham/tpl/debug"
	_ "github.com/gothamhq/gotham/tpl/encoding"
	_ "github.com/gothamhq/gotham/tpl/fmt"
	_ "github.com/gothamhq/gotham/tpl/gotham"
	_ "github.com/gothamhq/gotham/tpl/hugo"
	_ "github.com/gothamhq/gotham/tpl/images"
	_ "github.com/gothamhq/gotham/tpl/inflect"
	_ "github.com/gothamhq/gotham/tpl/js"
	_ "github.com/gothamhq/gotham/tpl/lang"
	_ "github.com/gothamhq/gotham/tpl/math"
	_ "github.com/gothamhq/gotham/tpl/openapi/openapi3"
	_ "github.com/gothamhq/gotham/tpl/os"
	_ "github.com/gothamhq/gotham/tpl/partials"
	_ "github.com/gothamhq/gotham/tpl/path"
	_ "github.com/gothamhq/gotham/tpl/reflect"
	_ "github.com/gothamhq/gotham/tpl/resources"
	_ "github.com/gothamhq/gotham/tpl/safe"
	_ "github.com/gothamhq/gotham/tpl/site"
	_ "github.com/gothamhq/gotham/tpl/strings"
	_ "github.com/gothamhq/gotham/tpl/templates"
	_ "github.com/gothamhq/gotham/tpl/time"
	_ "github.com/gothamhq/gotham/tpl/transform"
	_ "github.com/gothamhq/gotham/tpl/urls"
)

var (
	_    texttemplate.ExecHelper = (*templateExecHelper)(nil)
	zero reflect.Value
)

type templateExecHelper struct {
	running bool // whether we're in server mode.
	funcs   map[string]reflect.Value
}

func (t *templateExecHelper) GetFunc(tmpl texttemplate.Preparer, name string) (reflect.Value, bool) {
	if fn, found := t.funcs[name]; found {
		return fn, true
	}
	return zero, false
}

func (t *templateExecHelper) GetMapValue(tmpl texttemplate.Preparer, receiver, key reflect.Value) (reflect.Value, bool) {
	if params, ok := receiver.Interface().(maps.Params); ok {
		// Case insensitive.
		keystr := strings.ToLower(key.String())
		v, found := params[keystr]
		if !found {
			return zero, false
		}
		return reflect.ValueOf(v), true
	}

	v := receiver.MapIndex(key)

	return v, v.IsValid()
}

func (t *templateExecHelper) GetMethod(tmpl texttemplate.Preparer, receiver reflect.Value, name string) (method reflect.Value, firstArg reflect.Value) {
	if t.running {
		// This is a hot path and receiver.MethodByName really shows up in the benchmarks,
		// so we maintain a list of method names with that signature.
		switch name {
		case "GetPage", "Render":
			if info, ok := tmpl.(tpl.Info); ok {
				if m := receiver.MethodByName(name + "WithTemplateInfo"); m.IsValid() {
					return m, reflect.ValueOf(info)
				}
			}
		}
	}

	return receiver.MethodByName(name), zero
}

func newTemplateExecuter(d *deps.Deps) (texttemplate.Executer, map[string]reflect.Value) {
	funcs := createFuncMap(d)
	funcsv := make(map[string]reflect.Value)

	for k, v := range funcs {
		vv := reflect.ValueOf(v)
		funcsv[k] = vv
	}

	// Duplicate Go's internal funcs here for faster lookups.
	for k, v := range template.GoFuncs {
		if _, exists := funcsv[k]; !exists {
			vv, ok := v.(reflect.Value)
			if !ok {
				vv = reflect.ValueOf(v)
			}
			funcsv[k] = vv
		}
	}

	for k, v := range texttemplate.GoFuncs {
		if _, exists := funcsv[k]; !exists {
			funcsv[k] = v
		}
	}

	exeHelper := &templateExecHelper{
		running: d.Running,
		funcs:   funcsv,
	}

	return texttemplate.NewExecuter(
		exeHelper,
	), funcsv
}

func createFuncMap(d *deps.Deps) map[string]interface{} {
	funcMap := template.FuncMap{}

	// Merge the namespace funcs
	for _, nsf := range internal.TemplateFuncsNamespaceRegistry {
		ns := nsf(d)
		if _, exists := funcMap[ns.Name]; exists {
			panic(ns.Name + " is a duplicate template func")
		}
		funcMap[ns.Name] = ns.Context
		for _, mm := range ns.MethodMappings {
			for _, alias := range mm.Aliases {
				if _, exists := funcMap[alias]; exists {
					panic(alias + " is a duplicate template func")
				}
				funcMap[alias] = mm.Method
			}
		}
	}

	if d.OverloadedTemplateFuncs != nil {
		for k, v := range d.OverloadedTemplateFuncs {
			funcMap[k] = v
		}
	}

	return funcMap
}
