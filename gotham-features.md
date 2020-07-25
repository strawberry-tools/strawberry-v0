# Gotham Features

While we work on our own website and documentation, we'll list the features for Gotham here.

Gotham is built upon the strong foundation of [Hugo](https://gohugo.io).
99% of our features (and code for that matter) come from Hugo.
You can learn about all of Hugo's features (if you're unfamiliar) [here](https://gohugo.io/about/features/).


## Features Unique to Gotham

### Sitemap Page Exclusion

You can exclude pages from the built-in Sitemap by excluding it in a page's Front Matter.
It would look something like this:

```
---
title: "My Page"
description: "This is a page that is public but I don't want it in my sitemap.
date: "2020-06-13"
sitemap:
  exclude: true
---

My page.... blah blah blah.
```

While this setting can be done page-by-page to opt them out of the sitemap, you can also set this in your Gotham config file to opt-out all pages by default.
Then, you can manually opt-in pages.

### Menu Items Support Opening in New Tabs

If you have a menu item that you'd prefer to open in a new tab (or window, depending on browser), you can now do so via the Gotham config file.
Here's an example:

```yaml
#... gotham config

menu:
  main:
    - name: "Home"
      url: "/"
    - name: "Blog"
      url: "/blog/"
    - name: "Twitter"
      url: "https://twitter.com/GothamHQ_"
      newtab: true
```

In the example above, the homepage and the blog will open in the same tab when clicked, as expected.
The Twitter menu item will open in a new tab.
This is a common practice when linking to external URLs.

### Open Development Site in Browser

When running the development server locally (gotham server), you can now optionally open the dev URL in your default browser.
This can be done using the `--open` flag. For example:

```bash
gotham serve --open
```

### YouTube Shortcode

- The shortcode has a title parameter.
- The shortcode has start and stop parameters, measured in seconds.
