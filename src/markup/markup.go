// Copyright 2019 The Hugo Authors. All rights reserved.
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

package markup

import (
	"strings"

	"github.com/gothamhq/gotham/markup/highlight"

	"github.com/gothamhq/gotham/markup/markup_config"

	"github.com/gothamhq/gotham/markup/goldmark"

	"github.com/gothamhq/gotham/markup/org"

	"github.com/gothamhq/gotham/markup/asciidocext"
	"github.com/gothamhq/gotham/markup/blackfriday"
	"github.com/gothamhq/gotham/markup/converter"
	"github.com/gothamhq/gotham/markup/mmark"
	"github.com/gothamhq/gotham/markup/pandoc"
	"github.com/gothamhq/gotham/markup/rst"
)

func NewConverterProvider(cfg converter.ProviderConfig) (ConverterProvider, error) {
	converters := make(map[string]converter.Provider)

	markupConfig, err := markup_config.Decode(cfg.Cfg)
	if err != nil {
		return nil, err
	}

	if cfg.Highlight == nil {
		h := highlight.New(markupConfig.Highlight)
		cfg.Highlight = func(code, lang, optsStr string) (string, error) {
			return h.Highlight(code, lang, optsStr)
		}
	}

	cfg.MarkupConfig = markupConfig

	add := func(p converter.ProviderProvider, aliases ...string) error {
		c, err := p.New(cfg)
		if err != nil {
			return err
		}

		name := c.Name()

		aliases = append(aliases, name)

		if strings.EqualFold(name, cfg.MarkupConfig.DefaultMarkdownHandler) {
			aliases = append(aliases, "markdown")
		}

		addConverter(converters, c, aliases...)
		return nil
	}

	if err := add(goldmark.Provider); err != nil {
		return nil, err
	}
	if err := add(blackfriday.Provider); err != nil {
		return nil, err
	}
	if err := add(mmark.Provider); err != nil {
		return nil, err
	}
	if err := add(asciidocext.Provider, "ad", "adoc"); err != nil {
		return nil, err
	}
	if err := add(rst.Provider); err != nil {
		return nil, err
	}
	if err := add(pandoc.Provider, "pdc"); err != nil {
		return nil, err
	}
	if err := add(org.Provider); err != nil {
		return nil, err
	}

	return &converterRegistry{
		config:     cfg,
		converters: converters,
	}, nil
}

type ConverterProvider interface {
	Get(name string) converter.Provider
	//Default() converter.Provider
	GetMarkupConfig() markup_config.Config
	Highlight(code, lang, optsStr string) (string, error)
}

type converterRegistry struct {
	// Maps name (md, markdown, blackfriday etc.) to a converter provider.
	// Note that this is also used for aliasing, so the same converter
	// may be registered multiple times.
	// All names are lower case.
	converters map[string]converter.Provider

	config converter.ProviderConfig
}

func (r *converterRegistry) Get(name string) converter.Provider {
	return r.converters[strings.ToLower(name)]
}

func (r *converterRegistry) Highlight(code, lang, optsStr string) (string, error) {
	return r.config.Highlight(code, lang, optsStr)
}

func (r *converterRegistry) GetMarkupConfig() markup_config.Config {
	return r.config.MarkupConfig
}

func addConverter(m map[string]converter.Provider, c converter.Provider, aliases ...string) {
	for _, alias := range aliases {
		m[alias] = c
	}
}