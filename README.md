# Strawberry - the main ingredient of your JAM stack

Strawberry is a Static Site Generator that is extremely fast, efficient, and modular.
It is written in [Go][go-site] and is a soft fork of [Hugo][hugo-site].

[Website](https://www.StrawberrySSG.com) |
[Contribution Guide](CONTRIBUTING.md) |
[Twitter](https://twitter.com/StrawberrySSG)

[![Go Report Card](https://goreportcard.com/badge/github.com/strawberryssg/strawberry-v0)](https://goreportcard.com/report/github.com/strawberryssg/strawberry-v0)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/strawberryssg/strawberry-v0)
[![CircleCI](https://circleci.com/gh/strawberryssg/strawberry-v0.svg?style=shield)](https://circleci.com/gh/strawberryssg/strawberry-v0)
[![codecov](https://codecov.io/gh/strawberryssg/strawberry-v0/branch/master/graph/badge.svg)](https://codecov.io/gh/strawberryssg/strawberry-v0)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)
[![GoDoc](https://godoc.org/github.com/gohugoio/hugo?status.svg)](https://pkg.go.dev/github.com/gohugoio/hugo)

**Strawberry is a relatively new project forked from Hugo.
Please stay tuned while we continue to update this readme and convert things over.**

## Overview

Strawberry is a static HTML and CSS website generator written in [Go][].
It is optimized for speed, ease of use, and configurability.
Strawberry takes a directory with content and templates and renders them into a full HTML website.

Strawberry relies on Markdown files with front matter for metadata, and you can run Strawberry from any directory.
This works well for shared hosts and other systems where you don’t have a privileged account.

Strawberry renders a typical website of moderate size in a fraction of a second.
A good rule of thumb is that each piece of content renders in around 1 millisecond.

Strawberry is designed to work well for any kind of website including docs, blogs, marketing sites, newsletters, and more.

#### Supported Architectures

Strawberry is built for Linux, macOS, and Windows for amd64 (Intel and AMD) CPUs.
Support for ARM is planned.

**Complete documentation is available at [Strawberry Documentation](https://gohugo.io/getting-started/).**

## Choose How to Install

If you want to use Strawberry as your site generator, simply install the Strawberry binary.
The Strawberry binary have no external dependencies.

To contribute to the Strawberry source code or documentation, you should [fork the Strawberry GitHub project](https://github.com/strawberryssg/strawberry-v0#fork-destination-box) and clone it to your local machine.

Finally, you can install the Strawberry source code with `go`, build the binary yourself, and run Strawberry that way.
Building the binary is an easy task for an experienced `go` getter.

### Install Strawberry as Your Site Generator (Binary Install)

Use the [installation instructions in the Strawberry documentation](https://gohugo.io/getting-started/installing/).

### Build and Install the Binaries from Source (Advanced Install)

#### Prerequisite Tools

* [Git](https://git-scm.com/)
* [Go (we test it with the last 2 major versions; but note that Hugo 0.95.0 only builds with >= Go 1.18.)](https://golang.org/dl/)

#### Fetch from GitHub

Strawberry uses the Go Modules support built into Go 1.11 to build. The easiest is to clone Strawberry in a directory outside of `GOPATH`, as in the following example:

```bash
mkdir $HOME/src
cd $HOME/src
git clone https://github.com/strawberryssg/strawberry-v0.git
cd strawberry-v0
go install
```

**If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.**
	
## Contributing to Strawberry

For a complete guide to contributing to Strawberry, see the [Contribution Guide](CONTRIBUTING.md).

We welcome contributions to Strawberry of any kind including documentation, themes,
organization, tutorials, blog posts, bug reports, issues, feature requests,
feature implementations, pull requests, answering questions on the forum,
helping to manage issues, etc.

The Strawberry community and maintainers are [very active](https://github.com/strawberryssg/strawberry-v0/pulse/monthly) and helpful, and the project benefits greatly from this activity.

### Asking Support Questions

We're creating a forum soon where users and developers can ask questions.
You can use the GitHub Issue tracker to ask questions for now.

### Reporting Issues

If you believe you have found a defect in Strawberry or its documentation, use
the GitHub issue tracker to report the problem to the Strawberry maintainers.
When reporting the issue, please provide the version of Strawberry in use (`strawberry version`).

### Submitting Patches

The Strawberry project welcomes all contributors and contributions regardless of skill or experience level.
If you are interested in helping with the project, we will help you with your contribution.
Strawberry is a very active project with many contributions happening daily.

We want to create the best possible product for our users and the best contribution experience for our developers,
we have a set of guidelines which ensure that all contributions are acceptable.
The guidelines are not intended as a filter or barrier to participation.
If you are unfamiliar with the contribution process, the Strawberry team will help you and teach you how to bring your contribution in accordance with the guidelines.

For a complete guide to contributing code to Strawberry, see the [Contribution Guide](CONTRIBUTING.md).

[Go]: https://golang.org/
[Hugo Documentation]: https://gohugo.io/overview/introduction/

## Dependencies

Strawberry stands on the shoulder of many great open source libraries, especially [Hugo][hugo-site]:

If you run `strawberry env -v` you will get a complete and up to date list.

In Strawberry v0.20.0 that list is, in lexical order:

```
cloud.google.com/go/storage="v1.10.0"
cloud.google.com/go="v0.87.0"
github.com/Azure/azure-pipeline-go="v0.2.2"
github.com/Azure/azure-storage-blob-go="v0.9.0"
github.com/BurntSushi/locker="v0.0.0-20171006230638-a6e239ea1c69"
github.com/BurntSushi/toml="v0.3.1"
github.com/PuerkitoBio/purell="v1.1.1"
github.com/PuerkitoBio/urlesc="v0.0.0-20170810143723-de5bf2ad4578"
github.com/alecthomas/chroma="v0.9.4"
github.com/armon/go-radix="v1.0.0"
github.com/aws/aws-sdk-go="v1.41.14"
github.com/bep/debounce="v1.2.0"
github.com/bep/gitmap="v1.1.2"
github.com/bep/godartsass="v0.12.0"
github.com/bep/golibsass="v1.0.0"
github.com/bep/gowebp="v0.1.0"
github.com/bep/tmc="v0.5.1"
github.com/cli/safeexec="v1.0.0"
github.com/cpuguy83/go-md2man/v2="v2.0.0"
github.com/disintegration/gift="v1.2.1"
github.com/dlclark/regexp2="v1.4.0"
github.com/dustin/go-humanize="v1.0.0"
github.com/evanw/esbuild="v0.13.12"
github.com/fsnotify/fsnotify="v1.5.1"
github.com/getkin/kin-openapi="v0.80.0"
github.com/ghodss/yaml="v1.0.0"
github.com/go-openapi/jsonpointer="v0.19.5"
github.com/go-openapi/swag="v0.19.5"
github.com/gobuffalo/flect="v0.2.3"
github.com/gobwas/glob="v0.2.3"
github.com/gohugoio/go-i18n/v2="v2.1.3-0.20210430103248-4c28c89f8013"
github.com/gohugoio/locales="v0.14.0"
github.com/gohugoio/localescompressed="v0.14.0"
github.com/golang/groupcache="v0.0.0-20200121045136-8c9f03a8e57e"
github.com/golang/protobuf="v1.5.2"
github.com/google/go-cmp="v0.5.6"
github.com/google/uuid="v1.1.2"
github.com/google/wire="v0.4.0"
github.com/googleapis/gax-go/v2="v2.0.5"
github.com/googleapis/gax-go="v2.0.2+incompatible"
github.com/gorilla/websocket="v1.4.2"
github.com/inconshreveable/mousetrap="v1.0.0"
github.com/jdkato/prose="v1.2.1"
github.com/jmespath/go-jmespath="v0.4.0"
github.com/kyokomi/emoji/v2="v2.2.8"
github.com/mailru/easyjson="v0.0.0-20190626092158-b2ccc519800e"
github.com/mattn/go-ieproxy="v0.0.1"
github.com/mattn/go-isatty="v0.0.14"
github.com/mattn/go-runewidth="v0.0.9"
github.com/miekg/mmark="v1.3.6"
github.com/mitchellh/hashstructure="v1.1.0"
github.com/mitchellh/mapstructure="v1.4.2"
github.com/muesli/smartcrop="v0.3.0"
github.com/niklasfasching/go-org="v1.5.0"
github.com/olekukonko/tablewriter="v0.0.5"
github.com/pelletier/go-toml/v2="v2.0.0-beta.3.0.20210727221244-fa0796069526"
github.com/pkg/errors="v0.9.1"
github.com/rogpeppe/go-internal="v1.8.0"
github.com/russross/blackfriday/v2="v2.0.1"
github.com/russross/blackfriday="v1.5.3-0.20200218234912-41c5fccfd6f6"
github.com/rwcarlsen/goexif="v0.0.0-20190401172101-9e8deecbddbd"
github.com/sanity-io/litter="v1.5.1"
github.com/sass/libsass="3.6.5"
github.com/shurcooL/sanitized_anchor_name="v1.0.0"
github.com/spf13/afero="v1.6.0"
github.com/spf13/cast="v1.4.1"
github.com/spf13/cobra="v1.2.1"
github.com/spf13/fsync="v0.9.0"
github.com/spf13/jwalterweatherman="v1.1.0"
github.com/spf13/pflag="v1.0.5"
github.com/tdewolff/minify/v2="v2.9.22"
github.com/tdewolff/parse/v2="v2.5.21"
github.com/webmproject/libwebp="v1.2.0"
github.com/yuin/goldmark-highlighting="v0.0.0-20200307114337-60d527fdb691"
github.com/yuin/goldmark="v1.4.2"
go.opencensus.io="v0.23.0"
gocloud.dev="v0.20.0"
golang.org/x/image="v0.0.0-20210220032944-ac19c3e999fb"
golang.org/x/net="v0.0.0-20210614182718-04defd469f4e"
golang.org/x/oauth2="v0.0.0-20210628180205-a41e5a781914"
golang.org/x/sync="v0.0.0-20210220032951-036812b2e83c"
golang.org/x/sys="v0.0.0-20210908233432-aa78b53d3365"
golang.org/x/text="v0.3.7"
golang.org/x/xerrors="v0.0.0-20200804184101-5ec99f83aff1"
google.golang.org/api="v0.51.0"
google.golang.org/genproto="v0.0.0-20210716133855-ce7ef5c701ea"
google.golang.org/grpc="v1.39.0"
google.golang.org/protobuf="v1.27.1"
gopkg.in/yaml.v2="v2.4.0"
```


[go-site]: https://go.dev/
[hugo-site]: https://gohugo.io/
