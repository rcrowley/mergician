Mergician
=========

I like writing plain HTML. I like it so much more than using all the other tools I've used for making websites that I've been known to use `sed`(1) to make bulk edits or over-compromise on navigation to avoid having to use a higher-level CMS. Anyway, Mergician is the beginning of a truly minimalist CMS.

Usage
-----

```sh
mergician [-o <output>] <input>[...]
```

* `-o <output>`: write to this file instead of standard output
* `<input>[...]`:  pathname to one or more input HTML files

Example
-------

Consider these inputs:

`template.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<link href="template.css" rel="stylesheet">
<meta charset="utf-8">
<meta content="width=device-width,initial-scale=1" name="viewport">
<title>Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/>
<article class="body"></article>
<br/>
<footer><p>&copy; 2023</p></footer>
</body>
</html>
```

`article.html`:

```html
<!DOCTYPE html>
<html>
<head>
<link href="template.css" rel="stylesheet">
<title>My cool webpage</title>
</head>
<body>
<h1>Things</h1>
<p>Stuff</p>
</body>
</html>
```

`mergician template.html article.html` (so ordered to make `xargs`(1)-like invocations more natural) will combine the two and write the following to standard output:

```html
<!DOCTYPE html><html lang="en"><head>
<link href="template.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>Website / My cool webpage</title>
<link href="template.css" rel="stylesheet"/>
</head>
<body>
<header><h1>Website</h1></header>
<br/>
<article class="body">
<h1>Things</h1>
<p>Stuff</p>
</article>
<br/>
<footer><p>© 2023</p></footer>
</body></html>
```

Merge algorithm
---------------

Mergician merges the second HTML document it's given into the first, then the third into the first, and so on until all the documents are processed.

For each pair of documents it merges, it follows these steps:

1. Merge `<title>` tags by appending " / " and the second `<title>` tag's text to the first `<title>` tag's text. (TODO: If there is no first `<title>` tag, copy the second `<title>` tag in unchanged.)
2. Merge `<head>` tags by appending all the children from the second `<head>` tag into the first.
3. Merge `<body>` tags by appending all the children from the second `<body>` tag into the `<article class="body">`, `<div class="body">`, or `<section class="body">` in the first. (TODO: Extend this microformats-like structure as described below.)

Configuration
-------------

At present, the rules of the merge algorithm are fixed. However, I have a plan to expose the rules to reconfiguration to support other merge strategies and microformats. Here is how the default merge algorithm would be expressed in the configuration language:

```
<title> += " / " + <title>
<head> += <head> - <title>
<article class="body"> = <body>
<div class="body"> = <body>
<section class="body"> = <body>
```

I think there's probably value (especially for some of the imagined CMS use-cases below) to differentiate between merging in an element versus that element's children. For example, when merging HTML documents using this default configuration it's imperative to merge in the `<body>` tag's children since it's nonsense for a `<body>` tag to appear nested within another `<body>` tag, while when constructing a reverse-chronological index it might make more sense to keep the `<section class="index">` tags whole to provide some structural separation between each entry in the index. The syntax could look like this, which alludes to CSS (a plus) but also looks like a syntax error (when thinking in terms of r-values in most programming languages).

```
<article class="body"> = <body> *                  # merge <body>'s children
<section class="body"> += <section class="index"> # merge each <section class="index"> whole
```

CMS
---

The forthcoming CMS will build upon this merge algorithm to additionally support some or all of the features one might expect from a modern CMS.

* Reverse-chronological index
* Atom/RSS feeds
* `sitemap.xml`
* Lambda function invocation on publish (to e.g. send email to subscribers)

As a teaser, here are some configurations and invocations that might power parts of the CMS:

* Construct a reverse-chronological index:
    * Configuration: `<section class="index"> += <h1>`
    * Invocation: `mergician index.html newest.html ... oldest.html`
* Render to both HTML for the Web and an HTML email:
    * Web: `mergician article-template.html example.html`
    * Email: `mergician email-template.html example.html`

Don't let this imply that the CMS will be a shell program, though. I'm sure it will be a high-level Go program and it will use Mergician's configuration and merge algorithms heavily.
