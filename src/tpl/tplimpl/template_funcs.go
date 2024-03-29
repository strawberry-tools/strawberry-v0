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
	"context"
	"reflect"
	"strings"

	"github.com/strawberryssg/strawberry-v0/common/hreflect"
	"github.com/strawberryssg/strawberry-v0/common/maps"
	"github.com/strawberryssg/strawberry-v0/deps"
	"github.com/strawberryssg/strawberry-v0/tpl"
	"github.com/strawberryssg/strawberry-v0/tpl/internal"

	template "github.com/strawberryssg/strawberry-v0/tpl/internal/go_templates/htmltemplate"
	texttemplate "github.com/strawberryssg/strawberry-v0/tpl/internal/go_templates/texttemplate"

	// Init the namespaces
	_ "github.com/strawberryssg/strawberry-v0/tpl/cast"
	_ "github.com/strawberryssg/strawberry-v0/tpl/collections"
	_ "github.com/strawberryssg/strawberry-v0/tpl/compare"
	_ "github.com/strawberryssg/strawberry-v0/tpl/crypto"
	_ "github.com/strawberryssg/strawberry-v0/tpl/data"
	_ "github.com/strawberryssg/strawberry-v0/tpl/debug"
	_ "github.com/strawberryssg/strawberry-v0/tpl/diagrams"
	_ "github.com/strawberryssg/strawberry-v0/tpl/encoding"
	_ "github.com/strawberryssg/strawberry-v0/tpl/fmt"
	_ "github.com/strawberryssg/strawberry-v0/tpl/gotham"
	_ "github.com/strawberryssg/strawberry-v0/tpl/hugo"
	_ "github.com/strawberryssg/strawberry-v0/tpl/images"
	_ "github.com/strawberryssg/strawberry-v0/tpl/inflect"
	_ "github.com/strawberryssg/strawberry-v0/tpl/js"
	_ "github.com/strawberryssg/strawberry-v0/tpl/lang"
	_ "github.com/strawberryssg/strawberry-v0/tpl/math"
	_ "github.com/strawberryssg/strawberry-v0/tpl/openapi/openapi3"
	_ "github.com/strawberryssg/strawberry-v0/tpl/os"
	_ "github.com/strawberryssg/strawberry-v0/tpl/partials"
	_ "github.com/strawberryssg/strawberry-v0/tpl/path"
	_ "github.com/strawberryssg/strawberry-v0/tpl/reflect"
	_ "github.com/strawberryssg/strawberry-v0/tpl/resources"
	_ "github.com/strawberryssg/strawberry-v0/tpl/safe"
	_ "github.com/strawberryssg/strawberry-v0/tpl/site"
	_ "github.com/strawberryssg/strawberry-v0/tpl/strawberry"
	_ "github.com/strawberryssg/strawberry-v0/tpl/strings"
	_ "github.com/strawberryssg/strawberry-v0/tpl/templates"
	_ "github.com/strawberryssg/strawberry-v0/tpl/time"
	_ "github.com/strawberryssg/strawberry-v0/tpl/transform"
	_ "github.com/strawberryssg/strawberry-v0/tpl/urls"
)

var (
	_                texttemplate.ExecHelper = (*templateExecHelper)(nil)
	zero             reflect.Value
	contextInterface = reflect.TypeOf((*context.Context)(nil)).Elem()
)

type templateExecHelper struct {
	running bool // whether we're in server mode.
	funcs   map[string]reflect.Value
}

func (t *templateExecHelper) GetFunc(ctx context.Context, tmpl texttemplate.Preparer, name string) (fn reflect.Value, firstArg reflect.Value, found bool) {
	if fn, found := t.funcs[name]; found {
		if fn.Type().NumIn() > 0 {
			first := fn.Type().In(0)
			if first.Implements(contextInterface) {
				// TODO(bep) check if we can void this conversion every time -- and if that matters.
				// The first argument may be context.Context. This is never provided by the end user, but it's used to pass down
				// contextual information, e.g. the top level data context (e.g. Page).
				return fn, reflect.ValueOf(ctx), true
			}
		}

		return fn, zero, true
	}
	return zero, zero, false
}

func (t *templateExecHelper) Init(ctx context.Context, tmpl texttemplate.Preparer) {
}

func (t *templateExecHelper) GetMapValue(ctx context.Context, tmpl texttemplate.Preparer, receiver, key reflect.Value) (reflect.Value, bool) {
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

func (t *templateExecHelper) GetMethod(ctx context.Context, tmpl texttemplate.Preparer, receiver reflect.Value, name string) (method reflect.Value, firstArg reflect.Value) {
	if t.running {
		switch name {
		case "GetPage", "Render":
			if info, ok := tmpl.(tpl.Info); ok {
				if m := receiver.MethodByName(name + "WithTemplateInfo"); m.IsValid() {
					return m, reflect.ValueOf(info)
				}
			}
		}
	}

	fn := hreflect.GetMethodByName(receiver, name)
	if !fn.IsValid() {
		return zero, zero
	}

	if fn.Type().NumIn() > 0 {
		first := fn.Type().In(0)
		if first.Implements(contextInterface) {
			// The first argument may be context.Context. This is never provided by the end user, but it's used to pass down
			// contextual information, e.g. the top level data context (e.g. Page).
			return fn, reflect.ValueOf(ctx)
		}
	}

	return fn, zero
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

func createFuncMap(d *deps.Deps) map[string]any {
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
