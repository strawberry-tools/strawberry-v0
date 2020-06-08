# Gotham - an awesome static site generator

Gotham is a static site generator that is extremely fast, efficient, and modular.
It is written in [Go][go-site] and is a soft fork of [Hugo][hugo-site].

[Website](https://GothamHQ.com) |
[Contribution Guide](CONTRIBUTING.md) |
[Twitter](https://twitter.com/GothamHQ_)

[![Go Report Card](https://goreportcard.com/badge/github.com/gothamhq/gotham)](https://goreportcard.com/report/github.com/gothamhq/gotham)

**Gotham is a very new project only recently forked from Hugo.
Please stay tuned while we continue to update this readme and convert things over.**

## Overview

Gotham is a static HTML and CSS website generator written in [Go][].
It is optimized for speed, ease of use, and configurability.
Gotham takes a directory with content and templates and renders them into a full HTML website.

Gotham relies on Markdown files with front matter for metadata, and you can run Gotham from any directory.
This works well for shared hosts and other systems where you don’t have a privileged account.

Gotham renders a typical website of moderate size in a fraction of a second.
A good rule of thumb is that each piece of content renders in around 1 millisecond.

Gotham is designed to work well for any kind of website including docs, blogs, marketing sites, newsletters, and more.

#### Supported Architectures

Currently, we provide pre-built Gotham binaries for Windows, Linux, FreeBSD, NetBSD, DragonFly BSD, Open BSD, macOS (Darwin), and [Android](https://gist.github.com/bep/a0d8a26cf6b4f8bc992729b8e50b480b) for x64, i386 and ARM architectures.

Gotham may also be compiled from source wherever the Go compiler tool chain can run, e.g. for other operating systems including Plan 9 and Solaris.

**Complete documentation is available at [Gotham Documentation](https://gohugo.io/getting-started/).**

## Choose How to Install

If you want to use Gotham as your site generator, simply install the Gotham binary.
The Gotham binary have no external dependencies.

To contribute to the Gotham source code or documentation, you should [fork the Gotham GitHub project](https://github.com/gothamhq/gotham#fork-destination-box) and clone it to your local machine.

Finally, you can install the Gotham source code with `go`, build the binary yourself, and run Gotham that way.
Building the binary is an easy task for an experienced `go` getter.

### Install Gotham as Your Site Generator (Binary Install)

Use the [installation instructions in the Gotham documentation](https://gohugo.io/getting-started/installing/).

### Build and Install the Binaries from Source (Advanced Install)

#### Prerequisite Tools

* [Git](https://git-scm.com/)
* [Go (we test it with the last 2 major versions)](https://golang.org/dl/)

#### Fetch from GitHub

Gotham uses the Go Modules support built into Go 1.11 to build. The easiest is to clone Gotham in a directory outside of `GOPATH`, as in the following example:

```bash
mkdir $HOME/src
cd $HOME/src
git clone https://github.com/gothamhq/gotham.git
cd gotham
go install
```

**If you are a Windows user, substitute the `$HOME` environment variable above with `%USERPROFILE%`.**
	
## Contributing to Gotham

For a complete guide to contributing to Gotham, see the [Contribution Guide](CONTRIBUTING.md).

We welcome contributions to Gotham of any kind including documentation, themes,
organization, tutorials, blog posts, bug reports, issues, feature requests,
feature implementations, pull requests, answering questions on the forum,
helping to manage issues, etc.

The Gotham community and maintainers are [very active](https://github.com/gothamhq/gotham/pulse/monthly) and helpful, and the project benefits greatly from this activity.

### Asking Support Questions

We're creating a forum soon where users and developers can ask questions.
You can use the GitHub Issue tracker to ask questions for now.

### Reporting Issues

If you believe you have found a defect in Gotham or its documentation, use
the GitHub issue tracker to report the problem to the Gotham maintainers.
When reporting the issue, please provide the version of Gotham in use (`gotham version`).

### Submitting Patches

The Gotham project welcomes all contributors and contributions regardless of skill or experience level.
If you are interested in helping with the project, we will help you with your contribution.
Gotham is a very active project with many contributions happening daily.

We want to create the best possible product for our users and the best contribution experience for our developers,
we have a set of guidelines which ensure that all contributions are acceptable.
The guidelines are not intended as a filter or barrier to participation.
If you are unfamiliar with the contribution process, the Gotham team will help you and teach you how to bring your contribution in accordance with the guidelines.

For a complete guide to contributing code to Gotham, see the [Contribution Guide](CONTRIBUTING.md).


## Dependencies

Gotham stands on the shoulder of many great open source libraries, especially [Hugo][hugo-site]:

 | Dependency  | License |
 | :------------- | :------------- |
 | [github.com/gohugoio/hugo](https://github.com/gohugoio/hugo) | Apache License 2.0 |
 | [github.com/alecthomas/chroma](https://github.com/alecthomas/chroma) | MIT License |
 | [github.com/armon/go-radix](https://github.com/armon/go-radix) | MIT License |
 | [github.com/aws/aws-sdk-go](https://github.com/aws/aws-sdk-go) | Apache License 2.0 |
 | [github.com/bep/debounce](https://github.com/bep/debounce) | MIT License |
 | [github.com/bep/gitmap](https://github.com/bep/gitmap) | MIT License |
 | [github.com/bep/golibsass](https://github.com/bep/golibsass) | MIT License |
 | [github.com/bep/tmc](https://github.com/bep/tmc) | MIT License |
 | [github.com/BurntSushi/locker](https://github.com/BurntSushi/locker) | The Unlicense |
 | [github.com/BurntSushi/toml](https://github.com/BurntSushi/toml) | MIT License |
 | [github.com/cpuguy83/go-md2man](https://github.com/cpuguy83/go-md2man) | MIT License |
 | [github.com/danwakefield/fnmatch](https://github.com/danwakefield/fnmatch) | BSD 2-Clause "Simplified" License |
 | [github.com/disintegration/gift](https://github.com/disintegration/gift) | MIT License |
 | [github.com/dustin/go-humanize](https://github.com/dustin/go-humanize) | MIT License |
 | [github.com/fsnotify/fsnotify](https://github.com/fsnotify/fsnotify) | BSD 3-Clause "New" or "Revised" License |
 | [github.com/gobwas/glob](https://github.com/gobwas/glob) | MIT License |
 | [github.com/gorilla/websocket](https://github.com/gorilla/websocket) | BSD 2-Clause "Simplified" License |
 | [github.com/hashicorp/golang-lru](https://github.com/hashicorp/golang-lru) | Mozilla Public License 2.0 |
 | [github.com/hashicorp/hcl](https://github.com/hashicorp/hcl) | Mozilla Public License 2.0 |
 | [github.com/jdkato/prose](https://github.com/jdkato/prose) | MIT License |
 | [github.com/kr/pretty](https://github.com/kr/pretty) | MIT License |
 | [github.com/kyokomi/emoji](https://github.com/kyokomi/emoji) | MIT License |
 | [github.com/magiconair/properties](https://github.com/magiconair/properties) | BSD 2-Clause "Simplified" License |
 | [github.com/markbates/inflect](https://github.com/markbates/inflect) | MIT License |
 | [github.com/mattn/go-isatty](https://github.com/mattn/go-isatty) | MIT License |
 | [github.com/mattn/go-runewidth](https://github.com/mattn/go-runewidth) | MIT License |
 | [github.com/miekg/mmark](https://github.com/miekg/mmark) | Simplified BSD License |
 | [github.com/mitchellh/hashstructure](https://github.com/mitchellh/hashstructure) | MIT License |
 | [github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure) | MIT License |
 | [github.com/muesli/smartcrop](https://github.com/muesli/smartcrop) | MIT License |
 | [github.com/nicksnyder/go-i18n](https://github.com/nicksnyder/go-i18n) | MIT License |
 | [github.com/niklasfasching/go-org](https://github.com/niklasfasching/go-org) | MIT License |
 | [github.com/olekukonko/tablewriter](https://github.com/olekukonko/tablewriter) | MIT License |
 | [github.com/pelletier/go-toml](https://github.com/pelletier/go-toml) | MIT License |
 | [github.com/pkg/errors](https://github.com/pkg/errors) | BSD 2-Clause "Simplified" License |
 | [github.com/PuerkitoBio/purell](https://github.com/PuerkitoBio/purell) | BSD 3-Clause "New" or "Revised" License |
 | [github.com/PuerkitoBio/urlesc](https://github.com/PuerkitoBio/urlesc) | BSD 3-Clause "New" or "Revised" License |
 | [github.com/rogpeppe/go-internal](https://github.com/rogpeppe/go-internal) | BSD 3-Clause "New" or "Revised" License |
 | [github.com/russross/blackfriday](https://github.com/russross/blackfriday)  | Simplified BSD License |
 | [github.com/rwcarlsen/goexif](https://github.com/rwcarlsen/goexif) | BSD 2-Clause "Simplified" License |
 | [github.com/spf13/afero](https://github.com/spf13/afero) | Apache License 2.0 |
 | [github.com/spf13/cast](https://github.com/spf13/cast) | MIT License |
 | [github.com/spf13/cobra](https://github.com/spf13/cobra) | Apache License 2.0 |
 | [github.com/spf13/fsync](https://github.com/spf13/fsync) | MIT License |
 | [github.com/spf13/jwalterweatherman](https://github.com/spf13/jwalterweatherman) | MIT License |
 | [github.com/spf13/pflag](https://github.com/spf13/pflag) | BSD 3-Clause "New" or "Revised" License |
 | [github.com/spf13/viper](https://github.com/spf13/viper) | MIT License |
 | [github.com/tdewolff/minify](https://github.com/tdewolff/minify) | MIT License |
 | [github.com/tdewolff/parse](https://github.com/tdewolff/parse) | MIT License |
 | [github.com/yuin/goldmark](https://github.com/yuin/goldmark) | MIT License |
 | [github.com/yuin/goldmark-highlighting](https://github.com/yuin/goldmark-highlighting) | MIT License |
 | [go.opencensus.io](https://go.opencensus.io) | Apache License 2.0 |
 | [go.uber.org/atomic](https://go.uber.org/atomic) | MIT License |
 | [gocloud.dev](https://gocloud.dev) | Apache License 2.0 |
 | [golang.org/x/image](https://golang.org/x/image) | BSD 3-Clause "New" or "Revised" License |
 | [golang.org/x/net](https://golang.org/x/net) | BSD 3-Clause "New" or "Revised" License |
 | [golang.org/x/oauth2](https://golang.org/x/oauth2) | BSD 3-Clause "New" or "Revised" License |
 | [golang.org/x/sync](https://golang.org/x/sync) | BSD 3-Clause "New" or "Revised" License |
 | [golang.org/x/sys](https://golang.org/x/sys) | BSD 3-Clause "New" or "Revised" License |
 | [golang.org/x/text](https://golang.org/x/text) | BSD 3-Clause "New" or "Revised" License |
 | [golang.org/x/xerrors](https://golang.org/x/xerrors) | BSD 3-Clause "New" or "Revised" License |
 | [google.golang.org/api](https://google.golang.org/api) | BSD 3-Clause "New" or "Revised" License |
 | [google.golang.org/genproto](https://google.golang.org/genproto) | Apache License 2.0 |
 | [gopkg.in/ini.v1](https://gopkg.in/ini.v1) | Apache License 2.0 |
 | [gopkg.in/yaml.v2](https://gopkg.in/yaml.v2) | Apache License 2.0 |



[go-site]: https://go.dev/
[hugo-site]: https://gohugo.io/
