# Gotham Features

Gotham is built upon the strong foundation of [Hugo](https://gohugo.io).
99% of our features (and code for that matter) come from Hugo.
You can learn about all of Hugo's features (if you're unfamiliar) [here](https://gohugo.io/about/features/).

The current list of Gotham unique features can be found on [the website](https://www.gothamhq.com/features/).
Below are features tracked for the next or some future release.


## Features Unique to Gotham

### Apple App Site Association

A Gotham site can generate an associated domain file (`apple-app-site-association`) by providing two values in your Gotham config. 
This allows you connect your app and your Gotham website to provide both a native app and a browser experience.

```yaml
#... gotham config
aasaPrefix: "<my-application-identifier-prefix>"
aasaBundle: "<my-bundle-identifier>"
aasaVersion: 2
```

Gotham supports multiple versions of the Apple App Site Association file via the `aasaVersion` key.
The default value is set to `2`, which uses the latest version that Apple released in 2019. If you would like to use the earlier version you can pass `1` to the `aasaVersion` key. This key expects an integer.

Resources:

- https://developer.apple.com/documentation/safariservices/supporting_associated_domains
- https://search.developer.apple.com/appsearch-validation-tool

### JSON Feeds

Gotham pre-bakes [JSON Feeds](https://www.jsonfeed.org/) for you as it would RSS feeds.
The limit for the number of pages can be set in the Gotham config with the `jsonFeedLimit` key.
For example:

```yaml
jsonFeedLimit: 10
```

Would limit each feed to a max of 10 pages/items.
A value of `-1` means "all", which is the default.

JSONFeeds can be turned off completely by disabling the kind `JSONFeed`.
For example:

```yaml
disableKinds:
  - JSONFeed
```

This can also be done at runtime:

```bash
gotham --disableKinds=JSONFeed
```

By default, Gotham shows the full page content in a JSON Feed.
This functionality can be switched to just show the summary instead by setting `jsonFeedFull` to false.
Here's an example:

```yaml
jsonFeedFull: false
```
