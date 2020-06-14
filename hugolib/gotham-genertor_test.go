// Copyright 2020 Gotham Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package hugolib

// This file should be in the "transform/metainject" package because that's
// where the code it's testing lives. Unfortunately, the test site tools we
// need are in this "hugolib" package and are not exported. So for now, this
// test file will live here.

import (
	"fmt"
	"testing"

	"github.com/gothamhq/gotham/common/hugo"
)

// Test if the Gotham meta generator tag was injected into the homepage correctly
func TestGothamGeneratorInject(t *testing.T) {

	t.Parallel()

	siteConfig := `
baseurl = "http://example.com"
title = "Section Menu"

[[menu.main]]
    name    = "Home"
	url     = "/"
	weight  = -1
[[menu.main]]
    name    = "Blog"
	url     = "/blog/"
	newtab  = false
`

	htmlTemplate := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>HTML5 boilerplate – all you really need…</title>
	<link rel="stylesheet" href="css/style.css">
</head>
<body id="home">
	<h1>{{ .Title }}</h1>
	<p>{{ .Permalink }}</p>
</body>
</html>
`

	b := newTestSitesBuilder(t).WithConfigFile("toml", siteConfig)
	b.WithTemplatesAdded("layouts/index.html", htmlTemplate)
	b.Build(BuildCfg{})

	// Check if our Gotham meta gen tag is present
	metaGenTag := fmt.Sprintf(`<meta name="generator" content="Gotham %s" />`, hugo.GothamVersion)
	b.AssertHome(metaGenTag)

	// Make sure that the Hugo meta gen tag is not present
	metaGenTag = fmt.Sprintf(`<meta name="generator" content="Hugo %s" />`, hugo.CurrentVersion)
	b.AssertFileContentInvert("public/index.html", metaGenTag)
}
