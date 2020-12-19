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
```

Resources:

- https://developer.apple.com/documentation/safariservices/supporting_associated_domains
- https://search.developer.apple.com/appsearch-validation-tool
