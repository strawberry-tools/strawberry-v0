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

// This file is autogenerated.

// Package embedded defines the internal templates that Hugo provides.
package embedded

// EmbeddedTemplates represents all embedded templates.
var EmbeddedTemplates = [][2]string{
	{`_default/aasa_v1.json`, `{
	"applinks": {
		"apps": [],
		"details": [{
			"appID": "{{ printf "%s.%s" .Site.Config.Services.AASA.Prefix .Site.Config.Services.AASA.Bundle }}",
			"paths": ["*"]
		}]
	}
}
`},
	{`_default/aasa_v2.json`, `{
	"applinks": {
		"details": [{
			"appIDs": [ "{{ printf "%s.%s" .Site.Config.Services.AASA.Prefix .Site.Config.Services.AASA.Bundle }}" ],
			"components": [{
					"#": "no_universal_links",
					"exclude": true,
					"comment": "Matches any URL whose fragment equals no_universal_links and instructs the system not to open it as a universal link"
				},
				{
					"/": "*",
					"comment": "Matches any URL on the domain (wildcard)"
				}
			]
		}]
	}
}
`},
	{`_default/assetlinks.json`, `[{
  "relation": ["delegate_permission/common.handle_all_urls"],
  "target" : { "namespace": "android_app", "package_name": "{{ .Site.Config.Services.AssetLinks.PackageName }}",
               "sha256_cert_fingerprints": ["{{ .Site.Config.Services.AssetLinks.Fingerprint }}"] }
}]
`},
	{`_default/feed.json`, `{{- $thePage := . -}}
{{- if .IsHome -}}{{ $thePage = $.Site }}{{- end -}}
{{- $pages2 := $thePage.RegularPages -}}
{{- $limit2 := .Site.Config.Services.JSONFeed.Limit -}}
{{- if ge $limit2 1 -}}
{{- $pages2 = $pages2 | first $limit2 -}}
{{- end -}}
{
	"version": "https://jsonfeed.org/version/1.1",
	"title": "{{ if eq  .Title  .Site.Title }}{{ .Site.Title }}{{ else }}{{ with .Title }}{{ . }} on {{ end }}{{ .Site.Title }}{{ end }}",
	"description": "Recent content {{ if ne  .Title  .Site.Title }}{{ with .Title }}in {{ . }} {{ end }}{{ end }}on {{ .Site.Title }}",
	"home_page_url": "{{ .Site.BaseURL }}",
	{{ with .OutputFormats.Get "JSON" -}}
	"feed_url": "{{ .Permalink }}",
	{{ end -}}
	{{ with .Site.LanguageCode -}}
	"language": "{{ . }}",
	{{ end -}}
	{{ with $.Param "icon" -}}
	"icon": "{{ . | absURL }}",
	{{ end -}}
	{{ with $.Param "favicon" -}}
	"favicon": "{{ . | absURL }}",
	{{ end -}}
	{{ with .Site.Author.name -}}
	"authors": [
		{
			"name": "{{ . }}"{{ with $.Site.Author.url }},
			"url": "{{ . }}"{{ end }}{{ with $.Site.Author.avatar }},
			"avatar": "{{ . | absURL }}"{{ end }}
		}
	],
	{{ end -}}
	"items": [
	{{- range $index, $element2 := $pages2 }}
		{
			"title": {{ $element2.Title | jsonify }},
			"date_published": "{{ .Date.Format "2006-01-02T15:04:05Z07:00" }}",
			"date_modified": "{{ .Lastmod.Format "2006-01-02T15:04:05Z07:00" }}",
			"id": "{{ $element2.Permalink }}",
			"url": "{{ $element2.Permalink }}",
			{{ with .Params.author -}}
			"authors": [
				{
					"name": "{{ . }}"
				}
			],
			{{ end -}}
			"content_html": {{ if .Site.Config.Services.JSONFeed.Full }}{{ $element2.Content | jsonify }}{{ else }}{{ $element2.Summary | jsonify }}{{ end }}
		}{{- if ne (add $index 1) (len $pages2) -}},{{- end }}
	{{ end -}}
	]
}
`},
	{`_default/robots.txt`, `User-agent: *`},
	{`_default/rss.xml`, `{{- $pctx := . -}}
{{- if .IsHome -}}{{ $pctx = .Site }}{{- end -}}
{{- $pages := slice -}}
{{- if or $.IsHome $.IsSection -}}
{{- $pages = $pctx.RegularPages -}}
{{- else -}}
{{- $pages = $pctx.Pages -}}
{{- end -}}
{{- $limit := .Site.Config.Services.RSS.Limit -}}
{{- if ge $limit 1 -}}
{{- $pages = $pages | first $limit -}}
{{- end -}}
{{- printf "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>" | safeHTML }}
<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom">
  <channel>
    <title>{{ if eq  .Title  .Site.Title }}{{ .Site.Title }}{{ else }}{{ with .Title }}{{.}} on {{ end }}{{ .Site.Title }}{{ end }}</title>
    <link>{{ .Permalink }}</link>
    <description>Recent content {{ if ne  .Title  .Site.Title }}{{ with .Title }}in {{.}} {{ end }}{{ end }}on {{ .Site.Title }}</description>
    <generator>Strawberry -- strawberryssg.com</generator>{{ with .Site.LanguageCode }}
    <language>{{.}}</language>{{end}}{{ with .Site.Author.email }}
    <managingEditor>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</managingEditor>{{end}}{{ with .Site.Author.email }}
    <webMaster>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</webMaster>{{end}}{{ with .Site.Copyright }}
    <copyright>{{.}}</copyright>{{end}}{{ if not .Date.IsZero }}
    <lastBuildDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</lastBuildDate>{{ end }}
    {{- with .OutputFormats.Get "RSS" -}}
    {{ printf "<atom:link href=%q rel=\"self\" type=%q />" .Permalink .MediaType | safeHTML }}
    {{- end -}}
    {{ range $pages }}
    <item>
      <title>{{ .Title }}</title>
      <link>{{ .Permalink }}</link>
      <pubDate>{{ .Date.Format "Mon, 02 Jan 2006 15:04:05 -0700" | safeHTML }}</pubDate>
      {{ with .Site.Author.email }}<author>{{.}}{{ with $.Site.Author.name }} ({{.}}){{end}}</author>{{end}}
      <guid>{{ .Permalink }}</guid>
      <description>{{ .Summary | html }}</description>
    </item>
    {{ end }}
  </channel>
</rss>
`},
	{`_default/sitemap.xml`, `{{ printf "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>" | safeHTML }}
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
  xmlns:xhtml="http://www.w3.org/1999/xhtml">
  {{ range where .Pages ".Sitemap.Exclude" false }}
    {{- if .Permalink -}}
  <url>
    <loc>{{ .Permalink }}</loc>{{ if not .Lastmod.IsZero }}
    <lastmod>{{ safeHTML ( .Lastmod.Format "2006-01-02T15:04:05-07:00" ) }}</lastmod>{{ end }}{{ with .Sitemap.ChangeFreq }}
    <changefreq>{{ . }}</changefreq>{{ end }}{{ if ge .Sitemap.Priority 0.0 }}
    <priority>{{ .Sitemap.Priority }}</priority>{{ end }}{{ if .IsTranslated }}{{ range .Translations }}
    <xhtml:link
                rel="alternate"
                hreflang="{{ .Language.Lang }}"
                href="{{ .Permalink }}"
                />{{ end }}
    <xhtml:link
                rel="alternate"
                hreflang="{{ .Language.Lang }}"
                href="{{ .Permalink }}"
                />{{ end }}
  </url>
    {{- end -}}
  {{ end }}
</urlset>
`},
	{`_default/sitemapindex.xml`, `{{ printf "<?xml version=\"1.0\" encoding=\"utf-8\" standalone=\"yes\"?>" | safeHTML }}
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  {{ range . }}
  <sitemap>
    <loc>{{ .SitemapAbsURL }}</loc>
    {{ if not .LastChange.IsZero }}
      <lastmod>{{ .LastChange.Format "2006-01-02T15:04:05-07:00" | safeHTML }}</lastmod>
    {{ end }}
  </sitemap>
  {{ end }}
</sitemapindex>
`},
	{`_default/strawberry.json`, `{
	"strawberry_version": "{{ strawberry.Version }}",
	"hugo_version": "{{ strawberry.HugoVersion }}",
	"build_date": "{{ strawberry.BuildDate }}",
	"pages": "{{ len .Site.RegularPages }}"
}
`},
	{`alias.html`, `<!DOCTYPE html><html><head><title>{{ .Permalink }}</title><link rel="canonical" href="{{ .Permalink }}"/><meta name="robots" content="noindex"><meta charset="utf-8" /><meta http-equiv="refresh" content="0; url={{ .Permalink }}" /></head></html>`},
	{`disqus.html`, `{{- $pc := .Site.Config.Privacy.Disqus -}}
{{- if not $pc.Disable -}}
{{ if .Site.DisqusShortname }}<div id="disqus_thread"></div>
<script type="application/javascript">
    var disqus_config = function () {
    {{with .Params.disqus_identifier }}this.page.identifier = '{{ . }}';{{end}}
    {{with .Params.disqus_title }}this.page.title = '{{ . }}';{{end}}
    {{with .Params.disqus_url }}this.page.url = '{{ . | html  }}';{{end}}
    };
    (function() {
        if (["localhost", "127.0.0.1"].indexOf(window.location.hostname) != -1) {
            document.getElementById('disqus_thread').innerHTML = 'Disqus comments not available by default when the website is previewed locally.';
            return;
        }
        var d = document, s = d.createElement('script'); s.async = true;
        s.src = '//' + {{ .Site.DisqusShortname }} + '.disqus.com/embed.js';
        s.setAttribute('data-timestamp', +new Date());
        (d.head || d.body).appendChild(s);
    })();
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by Disqus.</a></noscript>
<a href="https://disqus.com" class="dsq-brlink">comments powered by <span class="logo-disqus">Disqus</span></a>{{end}}
{{- end -}}`},
	{`google_analytics.html`, `{{- $pc := .Site.Config.Privacy.GoogleAnalytics -}}
{{- if not $pc.Disable }}{{ with .Site.GoogleAnalytics -}}
{{ if hasPrefix . "G-"}}
<script async src="https://www.googletagmanager.com/gtag/js?id={{ . }}"></script>
<script>
{{ template "__ga_js_set_doNotTrack" $ }}
if (!doNotTrack) {
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());
	gtag('config', '{{ . }}', { 'anonymize_ip': {{- $pc.AnonymizeIP -}} });
}
</script>
{{ else if hasPrefix . "UA-" }}
<script type="application/javascript">
{{ template "__ga_js_set_doNotTrack" $ }}
if (!doNotTrack) {
	(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	})(window,document,'script','https://www.google-analytics.com/analytics.js','ga');
	{{- if $pc.UseSessionStorage }}
	if (window.sessionStorage) {
		var GA_SESSION_STORAGE_KEY = 'ga:clientId';
		ga('create', '{{ . }}', {
	    'storage': 'none',
	    'clientId': sessionStorage.getItem(GA_SESSION_STORAGE_KEY)
	   });
	   ga(function(tracker) {
	    sessionStorage.setItem(GA_SESSION_STORAGE_KEY, tracker.get('clientId'));
	   });
   }
	{{ else }}
	ga('create', '{{ . }}', 'auto');
	{{ end -}}
	{{ if $pc.AnonymizeIP }}ga('set', 'anonymizeIp', true);{{ end }}
	ga('send', 'pageview');
}
</script>
{{- end -}}
{{- end }}{{ end -}}

{{- define "__ga_js_set_doNotTrack" -}}{{/* This is also used in the async version. */}}
{{- $pc := .Site.Config.Privacy.GoogleAnalytics -}}
{{- if not $pc.RespectDoNotTrack -}}
var doNotTrack = false;
{{- else -}}
var dnt = (navigator.doNotTrack || window.doNotTrack || navigator.msDoNotTrack);
var doNotTrack = (dnt == "1" || dnt == "yes");
{{- end -}}
{{- end -}}`},
	{`google_analytics_async.html`, `{{- $pc := .Site.Config.Privacy.GoogleAnalytics -}}
{{- if not $pc.Disable -}}
{{ with .Site.GoogleAnalytics }}
<script type="application/javascript">
{{ template "__ga_js_set_doNotTrack" $ }}
if (!doNotTrack) {
	window.ga=window.ga||function(){(ga.q=ga.q||[]).push(arguments)};ga.l=+new Date;
	{{- if $pc.UseSessionStorage }}
	if (window.sessionStorage) {
		var GA_SESSION_STORAGE_KEY = 'ga:clientId';
		ga('create', '{{ . }}', {
	    'storage': 'none',
	    'clientId': sessionStorage.getItem(GA_SESSION_STORAGE_KEY)
	   });
	   ga(function(tracker) {
	    sessionStorage.setItem(GA_SESSION_STORAGE_KEY, tracker.get('clientId'));
	   });
   }
	{{ else }}
	ga('create', '{{ . }}', 'auto');
	{{ end -}}
	{{ if $pc.AnonymizeIP }}ga('set', 'anonymizeIp', true);{{ end }}
	ga('send', 'pageview');
}
</script>
<script async src='https://www.google-analytics.com/analytics.js'></script>
{{ end }}
{{- end -}}
`},
	{`google_news.html`, `{{ if .IsPage }}{{ with .Params.news_keywords }}
  <meta name="news_keywords" content="{{ range $i, $kw := first 10 . }}{{ if $i }},{{ end }}{{ $kw }}{{ end }}" />
{{ end }}{{ end }}`},
	{`opengraph.html`, `<meta property="og:title" content="{{ .Title }}" />
<meta property="og:description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end }}" />
<meta property="og:type" content="{{ if .IsPage }}article{{ else }}website{{ end }}" />
<meta property="og:url" content="{{ .Permalink }}" />

{{- with $.Params.images -}}
{{- range first 6 . }}<meta property="og:image" content="{{ . | absURL }}" />{{ end -}}
{{- else -}}
{{- $images := $.Resources.ByType "image" -}}
{{- $featured := $images.GetMatch "*feature*" -}}
{{- if not $featured }}{{ $featured = $images.GetMatch "{*cover*,*thumbnail*}" }}{{ end -}}
{{- with $featured -}}
<meta property="og:image" content="{{ $featured.Permalink }}"/>
{{- else -}}
{{- with $.Site.Params.images }}<meta property="og:image" content="{{ index . 0 | absURL }}"/>{{ end -}}
{{- end -}}
{{- end -}}

{{- if .IsPage }}
{{- $iso8601 := "2006-01-02T15:04:05-07:00" -}}
<meta property="article:section" content="{{ .Section }}" />
{{ with .PublishDate }}<meta property="article:published_time" {{ .Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />{{ end }}
{{ with .Lastmod }}<meta property="article:modified_time" {{ .Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />{{ end }}
{{- end -}}

{{- with .Params.audio }}<meta property="og:audio" content="{{ . }}" />{{ end }}
{{- with .Params.locale }}<meta property="og:locale" content="{{ . }}" />{{ end }}
{{- with .Site.Params.title }}<meta property="og:site_name" content="{{ . }}" />{{ end }}
{{- with .Params.videos }}{{- range . }}
<meta property="og:video" content="{{ . | absURL }}" />
{{ end }}{{ end }}

{{- /* If it is part of a series, link to related articles */}}
{{- $permalink := .Permalink }}
{{- $siteSeries := .Site.Taxonomies.series }}
{{ with .Params.series }}{{- range $name := . }}
  {{- $series := index $siteSeries ($name | urlize) }}
  {{- range $page := first 6 $series.Pages }}
    {{- if ne $page.Permalink $permalink }}<meta property="og:see_also" content="{{ $page.Permalink }}" />{{ end }}
  {{- end }}
{{ end }}{{ end }}

{{- /* Facebook Page Admin ID for Domain Insights */}}
{{- with .Site.Social.facebook_admin }}<meta property="fb:admins" content="{{ . }}" />{{ end }}
`},
	{`pagination.html`, `{{- $validFormats := slice "default" "terse" }}

{{- $msg1 := "When passing a map to the internal pagination template, one of the elements must be named 'page', and it must be set to the context of the current page." }}
{{- $msg2 := "The 'format' specified in the map passed to the internal pagination template is invalid. Valid choices are: %s." }}

{{- $page := . }}
{{- $format := "default" }}

{{- if reflect.IsMap . }}
  {{- with .page }}
    {{- $page = . }}
  {{- else }}
    {{- errorf $msg1 }}
  {{- end }}
  {{- with .format }}
    {{- $format = lower . }}
  {{- end }}
{{- end }}

{{- if in $validFormats $format }}
  {{- if gt $page.Paginator.TotalPages 1 }}
    <ul class="pagination pagination-{{ $format }}">
      {{- partial (printf "partials/inline/pagination/%s" $format) $page }}
    </ul>
  {{- end }}
{{- else }}
  {{- errorf $msg2 (delimit $validFormats ", ") }}
{{- end -}}

{{/* Format: default
{{/* --------------------------------------------------------------------- */}}
{{- define "partials/inline/pagination/default" }}
  {{- with .Paginator }}
    {{- $currentPageNumber := .PageNumber }}

    {{- with .First }}
      {{- if ne $currentPageNumber .PageNumber }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="First" class="page-link" role="button"><span aria-hidden="true">&laquo;&laquo;</span></a>
      </li>
      {{- else }}
      <li class="page-item disabled">
        <a href="#" aria-disabled="true" aria-label="First" class="page-link" role="button" tabindex="-1"><span aria-hidden="true">&laquo;&laquo;</span></a>
      </li>
      {{- end }}
    {{- end }}

    {{- with .Prev }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="Previous" class="page-link" role="button"><span aria-hidden="true">&laquo;</span></a>
      </li>
    {{- else }}
      <li class="page-item disabled">
        <a href="#" aria-disabled="true" aria-label="Previous" class="page-link" role="button" tabindex="-1"><span aria-hidden="true">&laquo;</span></a>
      </li>
    {{- end }}

    {{- $slots := 5 }}
    {{- $start := math.Max 1 (sub .PageNumber (math.Floor (div $slots 2))) }}
    {{- $end := math.Min .TotalPages (sub (add $start $slots) 1) }}
    {{- if lt (add (sub $end $start) 1) $slots }}
      {{- $start = math.Max 1 (add (sub $end $slots) 1) }}
    {{- end }}

    {{- range $k := seq $start $end }}
      {{- if eq $.Paginator.PageNumber $k }}
      <li class="page-item active">
        <a href="#" aria-current="page" aria-label="Page {{ $k }}" class="page-link" role="button">{{ $k }}</a>
      </li>
      {{- else }}
      <li class="page-item">
        <a href="{{ (index $.Paginator.Pagers (sub $k 1)).URL }}" aria-label="Page {{ $k }}" class="page-link" role="button">{{ $k }}</a>
      </li>
      {{- end }}
    {{- end }}

    {{- with .Next }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="Next" class="page-link" role="button"><span aria-hidden="true">&raquo;</span></a>
      </li>
    {{- else }}
      <li class="page-item disabled">
        <a href="#" aria-disabled="true" aria-label="Next" class="page-link" role="button" tabindex="-1"><span aria-hidden="true">&raquo;</span></a>
      </li>
    {{- end }}

    {{- with .Last }}
      {{- if ne $currentPageNumber .PageNumber }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="Last" class="page-link" role="button"><span aria-hidden="true">&raquo;&raquo;</span></a>
      </li>
      {{- else }}
      <li class="page-item disabled">
        <a href="#" aria-disabled="true" aria-label="Last" class="page-link" role="button" tabindex="-1"><span aria-hidden="true">&raquo;&raquo;</span></a>
      </li>
      {{- end }}
    {{- end }}
  {{- end }}
{{- end -}}

{{/* Format: terse
{{/* --------------------------------------------------------------------- */}}
{{- define "partials/inline/pagination/terse" }}
  {{- with .Paginator }}
    {{- $currentPageNumber := .PageNumber }}

    {{- with .First }}
      {{- if ne $currentPageNumber .PageNumber }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="First" class="page-link" role="button"><span aria-hidden="true">&laquo;&laquo;</span></a>
      </li>
      {{- end }}
    {{- end }}

    {{- with .Prev }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="Previous" class="page-link" role="button"><span aria-hidden="true">&laquo;</span></a>
      </li>
    {{- end }}

    {{- $slots := 3 }}
    {{- $start := math.Max 1 (sub .PageNumber (math.Floor (div $slots 2))) }}
    {{- $end := math.Min .TotalPages (sub (add $start $slots) 1) }}
    {{- if lt (add (sub $end $start) 1) $slots }}
      {{- $start = math.Max 1 (add (sub $end $slots) 1) }}
    {{- end }}

    {{- range $k := seq $start $end }}
      {{- if eq $.Paginator.PageNumber $k }}
      <li class="page-item active">
        <a href="#" aria-current="page" aria-label="Page {{ $k }}" class="page-link" role="button">{{ $k }}</a>
      </li>
      {{- else }}
      <li class="page-item">
        <a href="{{ (index $.Paginator.Pagers (sub $k 1)).URL }}" aria-label="Page {{ $k }}" class="page-link" role="button">{{ $k }}</a>
      </li>
      {{- end }}
    {{- end }}

    {{- with .Next }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="Next" class="page-link" role="button"><span aria-hidden="true">&raquo;</span></a>
      </li>
    {{- end }}

    {{- with .Last }}
      {{- if ne $currentPageNumber .PageNumber }}
      <li class="page-item">
        <a href="{{ .URL }}" aria-label="Last" class="page-link" role="button"><span aria-hidden="true">&raquo;&raquo;</span></a>
      </li>
      {{- end }}
    {{- end }}
  {{- end }}
{{- end -}}
`},
	{`schema.html`, `<meta itemprop="name" content="{{ .Title }}">
<meta itemprop="description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end }}">

{{- if .IsPage -}}
{{- $iso8601 := "2006-01-02T15:04:05-07:00" -}}
{{ with .PublishDate }}<meta itemprop="datePublished" {{ .Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />{{ end}}
{{ with .Lastmod }}<meta itemprop="dateModified" {{ .Format $iso8601 | printf "content=%q" | safeHTMLAttr }} />{{ end}}
<meta itemprop="wordCount" content="{{ .WordCount }}">

{{- with $.Params.images -}}
{{- range first 6 . -}}<meta itemprop="image" content="{{ . | absURL }}">{{ end -}}
{{- else -}}
{{- $images := $.Resources.ByType "image" -}}
{{- $featured := $images.GetMatch "*feature*" -}}
{{- if not $featured }}{{ $featured = $images.GetMatch "{*cover*,*thumbnail*}" }}{{ end -}}
{{- with $featured -}}
<meta itemprop="image" content="{{ $featured.Permalink }}">
{{- else -}}
{{- with $.Site.Params.images -}}<meta itemprop="image" content="{{ index . 0 | absURL }}"/>{{ end -}}
{{- end -}}
{{- end -}}

<!-- Output all taxonomies as schema.org keywords -->
<meta itemprop="keywords" content="{{ if .IsPage}}{{ range $index, $tag := .Params.tags }}{{ $tag }},{{ end }}{{ else }}{{ range $plural, $terms := .Site.Taxonomies }}{{ range $term, $val := $terms }}{{ printf "%s," $term }}{{ end }}{{ end }}{{ end }}" />
{{- end -}}
`},
	{`shortcodes/__h_simple_assets.html`, `{{ define "__h_simple_css" }}{{/* These template definitions are global. */}}
{{- if not (.Page.Scratch.Get "__h_simple_css") -}}
{{/* Only include once */}}
{{-  .Page.Scratch.Set "__h_simple_css" true -}}
<style>
.__h_video {
   position: relative;
   padding-bottom: 56.23%;
   height: 0;
   overflow: hidden;
   width: 100%;
   background: #000;
}
.__h_video img {
   width: 100%;
   height: auto;
   color: #000;
}
.__h_video .play {
   height: 72px;
   width: 72px;
   left: 50%;
   top: 50%;
   margin-left: -36px;
   margin-top: -36px;
   position: absolute;
   cursor: pointer;
}
</style>
{{- end -}}
{{- end -}}
{{- define "__h_simple_icon_play" -}}
<svg version="1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 61 61"><circle cx="30.5" cy="30.5" r="30.5" opacity=".8" fill="#000"></circle><path d="M25.3 19.2c-2.1-1.2-3.8-.2-3.8 2.2v18.1c0 2.4 1.7 3.4 3.8 2.2l16.6-9.1c2.1-1.2 2.1-3.2 0-4.4l-16.6-9z" fill="#fff"></path></svg>
{{- end -}}
`},
	{`shortcodes/figure.html`, `<figure{{ with .Get "class" }} class="{{ . }}"{{ end }}>
    {{- if .Get "link" -}}
        <a href="{{ .Get "link" }}"{{ with .Get "target" }} target="{{ . }}"{{ end }}{{ with .Get "rel" }} rel="{{ . }}"{{ end }}>
    {{- end -}}
    <img src="{{ .Get "src" }}"
         {{- if or (.Get "alt") (.Get "caption") }}
         alt="{{ with .Get "alt" }}{{ . }}{{ else }}{{ .Get "caption" | markdownify| plainify }}{{ end }}"
         {{- end -}}
         {{- with .Get "width" }} width="{{ . }}"{{ end -}}
         {{- with .Get "height" }} height="{{ . }}"{{ end -}}
    /><!-- Closing img tag -->
    {{- if .Get "link" }}</a>{{ end -}}
    {{- if or (or (.Get "title") (.Get "caption")) (.Get "attr") -}}
        <figcaption>
            {{ with (.Get "title") -}}
                <h4>{{ . }}</h4>
            {{- end -}}
            {{- if or (.Get "caption") (.Get "attr") -}}<p>
                {{- .Get "caption" | markdownify -}}
                {{- with .Get "attrlink" }}
                    <a href="{{ . }}">
                {{- end -}}
                {{- .Get "attr" | markdownify -}}
                {{- if .Get "attrlink" }}</a>{{ end }}</p>
            {{- end }}
        </figcaption>
    {{- end }}
</figure>
`},
	{`shortcodes/gist.html`, `<script type="application/javascript" src="https://gist.github.com/{{ index .Params 0 }}/{{ index .Params 1 }}.js{{if len .Params | eq 3 }}?file={{ index .Params 2 }}{{end}}"></script>
`},
	{`shortcodes/highlight.html`, `{{ if len .Params | eq 2 }}{{ highlight (trim .Inner "\n\r") (.Get 0) (.Get 1) }}{{ else }}{{ highlight (trim .Inner "\n\r") (.Get 0) "" }}{{ end }}`},
	{`shortcodes/instagram.html`, `{{- $pc := site.Config.Privacy.Instagram -}}
{{- if not $pc.Disable -}}
  {{ $accessToken := site.Config.Services.Instagram.AccessToken }}
  {{- if not $accessToken -}}
    {{- erroridf "error-missing-instagram-accesstoken" "instagram shortcode: Missing config value for services.instagram.accessToken. This can be set in config.toml, but it is recommended to configure this via the HUGO_SERVICES_INSTAGRAM_ACCESSTOKEN OS environment variable. If you are using a Client Access Token, remember that you must combine it with your App ID using a pipe symbol (<APPID>|<CLIENTTOKEN>) otherwise the request will fail." -}}
  {{- else -}}
    {{- if $pc.Simple -}}
      {{ template "_internal/shortcodes/instagram_simple.html" . }}
    {{- else -}}
      {{ $id := .Get 0 }}
      {{ $hideCaption := cond (eq (.Get 1) "hidecaption") "1" "0" }}
      {{ $headers := dict "Authorization" (printf "Bearer %s" $accessToken) }}
      {{ with getJSON "https://graph.facebook.com/v8.0/instagram_oembed/?url=https://instagram.com/p/" $id "/&hidecaption=" $hideCaption $headers }}
        {{ .html | safeHTML }}
      {{ end }}
    {{- end -}}
  {{- end -}}
{{- end -}}`},
	{`shortcodes/instagram_simple.html`, `{{- $pc := .Page.Site.Config.Privacy.Instagram -}}
{{- $sc := .Page.Site.Config.Services.Instagram -}}
{{- if not $pc.Disable -}}
  {{ $accessToken := site.Config.Services.Instagram.AccessToken }}
  {{- if not $accessToken -}}
    {{- erroridf "error-missing-instagram-accesstoken" "instagram shortcode: Missing config value for services.instagram.accessToken. This can be set in config.toml, but it is recommended to configure this via the HUGO_SERVICES_INSTAGRAM_ACCESSTOKEN OS environment variable. If you are using a Client Access Token, remember that you must combine it with your App ID using a pipe symbol (<APPID>|<CLIENTTOKEN>) otherwise the request will fail." -}}
  {{- else -}}
    {{- $id := .Get 0 -}}
    {{- $headers := dict "Authorization" (printf "Bearer %s" $accessToken) -}}
    {{- $item := getJSON "https://graph.facebook.com/v8.0/instagram_oembed/?url=https://instagram.com/p/" $id "/&amp;maxwidth=640&amp;omitscript=true" $headers -}}
    {{- $class1 := "__h_instagram" -}}
    {{- $class2 := "s_instagram_simple" -}}
    {{- $hideCaption := (eq (.Get 1) "hidecaption") -}}
    {{ with $item }}
      {{- $mediaURL := printf "https://instagram.com/p/%s/" $id | safeURL -}}
      {{- if not $sc.DisableInlineCSS -}}
        {{ template "__h_simple_instagram_css" $ }}
      {{- end -}}
      <div class="{{ $class1 }} {{ $class2 }} card" style="max-width: {{ $item.thumbnail_width }}px">
        <div class="card-header">
          <a href="{{ $item.author_url | safeURL }}" class="card-link">
            {{ $item.author_name }}
          </a>
        </div>
        <a href="{{ $mediaURL }}" rel="noopener" target="_blank">
          <img class="card-img-top img-fluid" src="{{ $item.thumbnail_url }}" width="{{ $item.thumbnail_width }}"  height="{{ $item.thumbnail_height }}" alt="Instagram Image">
        </a>
        <div class="card-body">
          {{ if not $hideCaption }}
            <p class="card-text">
              <a href="{{ $item.author_url | safeURL }}" class="card-link">
                {{ $item.author_name }}
              </a>
              {{ $item.title}}
            </p>
          {{ end }}
          <a href="{{ $item.author_url | safeURL }}" class="card-link">
            View More on Instagram
          </a>
        </div>
      </div>
    {{ end }}
  {{- end -}}
{{- end -}}

{{ define "__h_simple_instagram_css" }}
  {{ if not (.Page.Scratch.Get "__h_simple_instagram_css") }}
    {{/* Only include once */}}
    {{  .Page.Scratch.Set "__h_simple_instagram_css" true }}
    <style type="text/css">
      .__h_instagram.card {
      font-family: -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif;
      font-size: 14px;
      border: 1px solid rgb(219, 219, 219);
      padding: 0;
      margin-top: 30px;
      }
      .__h_instagram.card .card-header, .__h_instagram.card .card-body {
      padding: 10px 10px 10px;
      }
      .__h_instagram.card img {
      width: 100%;
      	height: auto;
      }
    </style>
  {{ end }}
{{ end }}`},
	{`shortcodes/mastodon.html`, `{{ $url := .Get "url" | urls.Parse }}
{{ $api := printf "%s://%s/api/oembed/?" $url.Scheme $url.Host }}
{{ $width := .Get "width" | default "400" }}
{{ $height := .Get "height" | default "null" }}
{{ with getJSON $api (querify "url" $url "width" $width "height" $height) }}{{ .html | safeHTML }}{{ end }}
`},
	{`shortcodes/param.html`, `{{- $name := (.Get 0) -}}
{{- with $name -}}
{{- with ($.Page.Param .) }}{{ . }}{{ else }}{{ errorf "Param %q not found: %s" $name $.Position }}{{ end -}}
{{- else }}{{ errorf "Missing param key: %s" $.Position }}{{ end -}}`},
	{`shortcodes/qrcode.html`, `<figure{{ with .Get "class" }} class="{{ . }}"{{ end }}>
    <img {{ printf "src=%q" (qrcoder (.Get "url")) | safeHTMLAttr }}
         {{- if or (.Get "alt") (.Get "caption") }}
         alt="{{ with .Get "alt" }}{{ . }}{{ else }}{{ .Get "caption" | markdownify| plainify }}{{ end }}"
         {{- end -}}
		 {{- with .Get "size" }} width="{{ . }}" height="{{ . }}"{{- else }} width="256" height="256"{{ end }} /> <!-- Closing img tag -->
    {{- if or (or (.Get "title") (.Get "caption")) (.Get "attr") -}}
        <figcaption>
            {{ with (.Get "title") -}}
                <h4>{{ . }}</h4>
            {{- end -}}
            {{- if or (.Get "caption") (.Get "attr") -}}<p>
                {{- .Get "caption" | markdownify -}}
                {{- with .Get "attrlink" }}
                    <a href="{{ . }}">
                {{- end -}}
                {{- .Get "attr" | markdownify -}}
                {{- if .Get "attrlink" }}</a>{{ end }}</p>
            {{- end }}
        </figcaption>
    {{- end }}
</figure>
`},
	{`shortcodes/ref.html`, `{{ ref . .Params }}`},
	{`shortcodes/relref.html`, `{{ relref . .Params }}`},
	{`shortcodes/twitter.html`, `{{- $pc := .Page.Site.Config.Privacy.Twitter -}}
{{- if not $pc.Disable -}}
  {{- if $pc.Simple -}}
    {{- template "_internal/shortcodes/twitter_simple.html" . -}}
  {{- else -}}
    {{- $msg1 := "The %q shortcode requires two named parameters: user and id. See %s" -}}
    {{- $msg2 := "The %q shortcode will soon require two named parameters: user and id. See %s" -}}
    {{- if .IsNamedParams -}}
      {{- $id := .Get "id" -}}
      {{- $user := .Get "user" -}}
      {{- if and $id $user -}}
        {{- template "render-tweet" (dict "id" $id "user" $user "dnt" $pc.EnableDNT) -}}
      {{- else -}}
        {{- errorf $msg1 .Name .Position -}}
      {{- end -}}
    {{- else -}}
      {{- $id := .Get 1 -}}
      {{- $user := .Get 0 -}}
      {{- if eq 1 (len .Params) -}}
        {{- $id = .Get 0 -}}
        {{- $user = "x" -}} {{/* This triggers a redirect. It works, but may not work forever. */}}
        {{- warnf $msg2 .Name .Position -}}
      {{- end -}}
      {{- template "render-tweet" (dict "id" $id "user" $user "dnt" $pc.EnableDNT) -}}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{- define "render-tweet" -}}
  {{- $url := printf "https://twitter.com/%v/status/%v" .user .id -}}
  {{- $query := querify "url" $url "dnt" .dnt -}}
  {{- $request := printf "https://publish.twitter.com/oembed?%s" $query -}}
  {{- $json := getJSON $request -}}
  {{- $json.html | safeHTML -}}
{{- end -}}
`},
	{`shortcodes/twitter_simple.html`, `{{- $pc := .Page.Site.Config.Privacy.Twitter -}}
{{- $sc := .Page.Site.Config.Services.Twitter -}}
{{- if not $pc.Disable -}}
  {{- $msg1 := "The %q shortcode requires two named parameters: user and id. See %s" -}}
  {{- $msg2 := "The %q shortcode will soon require two named parameters: user and id. See %s" -}}
  {{- if .IsNamedParams -}}
    {{- $id := .Get "id" -}}
    {{- $user := .Get "user" -}}
    {{- if and $id $user -}}
      {{- template "render-simple-tweet" (dict "id" $id "user" $user "dnt" $pc.EnableDNT "disableInlineCSS" $sc.DisableInlineCSS "ctx" .) -}}
    {{- else -}}
      {{- errorf $msg1 .Name .Position -}}
    {{- end -}}
  {{- else -}}
    {{- $id := .Get 1 -}}
    {{- $user := .Get 0 -}}
    {{- if eq 1 (len .Params) -}}
      {{- $id = .Get 0 -}}
      {{- $user = "x" -}} {{/* This triggers a redirect. It works, but may not work forever. */}}
      {{- warnf $msg2 .Name .Position -}}
    {{- end -}}
    {{- template "render-simple-tweet" (dict "id" $id "user" $user "dnt" $pc.EnableDNT "disableInlineCSS" $sc.DisableInlineCSS "ctx" .) -}}
  {{- end -}}
{{- end -}}

{{- define "render-simple-tweet" -}}
  {{- $url := printf "https://twitter.com/%v/status/%v" .user .id -}}
  {{- $query := querify "url" $url "dnt" .dnt "omit_script" true -}}
  {{- $request := printf "https://publish.twitter.com/oembed?%s" $query -}}
  {{- $json := getJSON $request -}}
  {{- if not .disableInlineCSS -}}
    {{- template "__h_simple_twitter_css" .ctx -}}
  {{- end }}
  {{ $json.html | safeHTML -}}
{{- end -}}

{{- define "__h_simple_twitter_css" -}}
  {{- if not (.Page.Scratch.Get "__h_simple_twitter_css") -}}
    {{/* Only include once */}}
    {{- .Page.Scratch.Set "__h_simple_twitter_css" true }}
    <style type="text/css">
      .twitter-tweet {
        font: 14px/1.45 -apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Oxygen-Sans,Ubuntu,Cantarell,"Helvetica Neue",sans-serif;
        border-left: 4px solid #2b7bb9;
        padding-left: 1.5em;
        color: #555;
      }
      .twitter-tweet a {
        color: #2b7bb9;
        text-decoration: none;
      }
      blockquote.twitter-tweet a:hover,
      blockquote.twitter-tweet a:focus {
        text-decoration: underline;
      }
    </style>
  {{- end -}}
{{- end -}}
`},
	{`shortcodes/vimeo.html`, `{{- $pc := .Page.Site.Config.Privacy.Vimeo -}}
{{- if not $pc.Disable -}}
{{- if $pc.Simple -}}
{{ template "_internal/shortcodes/vimeo_simple.html" . }}
{{- else -}}
{{ if .IsNamedParams }}<div {{ if .Get "class" }}class="{{ .Get "class" }}"{{ else }}style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="https://player.vimeo.com/video/{{ .Get "id" }}{{- if $pc.EnableDNT -}}?dnt=1{{- end -}}" {{ if not (.Get "class") }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" {{ end }}{{ if .Get "title"}}title="{{ .Get "title" }}"{{ else }}title="vimeo video"{{ end }} webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</div>{{ else }}
<div {{ if gt (len .Params) 1 }}class="{{ .Get 1 }}"{{ else }}style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden;"{{ end }}>
  <iframe src="https://player.vimeo.com/video/{{ .Get 0 }}{{- if $pc.EnableDNT -}}?dnt=1{{- end -}}" {{ if len .Params | eq 1 }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" {{ end }}{{ if len .Params | eq 3 }}title="{{ .Get 2 }}"{{ else }}title="vimeo video"{{ end }} webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>
</div>
{{ end }}
{{- end -}}
{{- end -}}`},
	{`shortcodes/vimeo_simple.html`, `{{- $pc := .Page.Site.Config.Privacy.Vimeo -}}
{{- if not $pc.Disable -}}
{{ $id := .Get "id" | default (.Get 0) }}
{{ $dnt := cond (eq $pc.EnableDNT true) "?dnt=1" "" }}
{{- $item := getJSON (print "https://vimeo.com/api/oembed.json?url=https://vimeo.com/" $id $dnt) -}}
{{ $class := .Get "class" | default (.Get 1) }}
{{ $hasClass := $class }}
{{ $class := $class | default "__h_video" }}
{{ if not $hasClass }}
{{/* If class is set, assume the user wants to provide his own styles. */}}
{{ template "__h_simple_css" $ }}
{{ end }}
{{ $secondClass := "s_video_simple" }}
<div class="{{ $secondClass }} {{ $class }}">
{{- with $item }}
<a href="{{ .provider_url }}{{ .video_id }}" rel="noopener" target="_blank">
{{ $thumb := .thumbnail_url }}
{{ $original := $thumb | replaceRE "(_.*\\.)" "." }}
<img src="{{ $thumb }}" srcset="{{ $thumb }} 1x, {{ $original }} 2x" alt="{{ .title }}">
<div class="play">{{ template "__h_simple_icon_play" $ }}</div></a></div>
{{- end -}}
{{- end -}}`},
	{`shortcodes/youtube.html`, `{{- $pc := .Page.Site.Config.Privacy.YouTube -}}
{{- if not $pc.Disable -}}
{{- $ytHost := cond $pc.PrivacyEnhanced  "www.youtube-nocookie.com" "www.youtube.com" -}}
{{- $id := .Get "id" | default (.Get 0) -}}
{{- $class := .Get "class" | default (.Get 1) -}}
{{- $title := .Get "title" | default "YouTube Video" }}
{{ $urlParams := slice }}
{{ with .Get "autoplay" }}{{ $urlParams = $urlParams | append "autoplay=1" }}{{ end }}
{{ with .Get "start" }}{{ $urlParams = $urlParams | append (print "start=" .) }}{{ end }}
{{ with .Get "end" }}{{ $urlParams = $urlParams | append (print "end=" .) }}{{ end }}
{{ $urlParamsStr := delimit $urlParams "&" }}
<div {{ with $class }}class="{{ . }}"{{ else }}style="position: relative; padding-bottom: 56.25%; height: 0; overflow: hidden;"{{ end }}>
	<iframe src="https://{{ $ytHost }}/embed/{{ $id }}{{ if gt (len $urlParams) 0 }}?{{ $urlParamsStr | safeURL }}{{ end }}" {{ if not $class }}style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; border:0;" {{ end }}allowfullscreen title="{{ $title }}"></iframe>
</div>
{{ end -}}
`},
	{`twitter_cards.html`, `{{- with $.Params.images -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ index . 0 | absURL }}"/>
{{ else -}}
{{- $images := $.Resources.ByType "image" -}}
{{- $featured := $images.GetMatch "*feature*" -}}
{{- if not $featured }}{{ $featured = $images.GetMatch "{*cover*,*thumbnail*}" }}{{ end -}}
{{- with $featured -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ $featured.Permalink }}"/>
{{- else -}}
{{- with $.Site.Params.images -}}
<meta name="twitter:card" content="summary_large_image"/>
<meta name="twitter:image" content="{{ index . 0 | absURL }}"/>
{{ else -}}
<meta name="twitter:card" content="summary"/>
{{- end -}}
{{- end -}}
{{- end }}
<meta name="twitter:title" content="{{ .Title }}"/>
<meta name="twitter:description" content="{{ with .Description }}{{ . }}{{ else }}{{if .IsPage}}{{ .Summary }}{{ else }}{{ with .Site.Params.description }}{{ . }}{{ end }}{{ end }}{{ end -}}"/>
{{ with .Site.Social.twitter -}}
<meta name="twitter:site" content="@{{ . }}"/>
{{ end -}}
`},
}
