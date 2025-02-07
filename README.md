Mergician
=========

HTML is the language of the Web but the Web has moved beyond the bare, unstyled Web of the '90s (and university professors). Visitors to your site expect coherent layouts, visual consistency, navigation. This is the point at which most folks reach for a templating language. Mergician reimagines the messy mix of HTML and templating language. All you need are HTML documents. Instead of rendering a template, merge a content HTML document into a layout HTML document.

For example, if you define two files:

`layout.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<link href="layout.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/>
<div class="body"></div>
<br/>
<footer><p>&copy; 2025</p></footer>
</body>
</html>
```

`webpage.html`:

```html
<!DOCTYPE html>
<html>
<head>
<title>My cool webpage</title>
</head>
<body>
<h1>Things</h1>
<p>Stuff</p>
</body>
</html>
```

And use Mergician to merge them together:

```sh
mergician layout.html webpage.html
```

You get your webpage wrapped in your layout:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<link href="layout.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>My cool webpage — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/>
<div class="body">
<h1>Things</h1>
<p>Stuff</p>
</div>
<br/>
<footer><p>© 2025</p></footer>
</body>
</html>
```

Mergician uses Microformats-inspired (remember those?!) [rules](#rules) to control how HTML documents are merged. Every input is a complete, standalone HTML document that can be viewed and edited on its own, no template language artifacts anywhere.

Installation
------------

```sh
go install github.com/rcrowley/mergician@latest
```

Usage
-----

```sh
mergician [-o <output>] [-r <rule>[...]] <input>[...]
```

* `-o <output>`: write to this file instead of standard output
* `-r <rule>`: use a custom rule for merging inputs (overrides all defaults; may be repeated); each rule is a destination HTML tag with optional attributes, "=" or "+=", and a source HTML tag with optional attributes; the default rules are:
    * `<article class="body"> = <body>`
    * `<div class="body"> = <body>`
    * `<section class="body"> = <body>`
* `<input>[...]`: one or more input HTML, Markdown, or Google Doc HTML-in-zip documents

Rules
-----

Mergician rules control how two HTML documents are merged. Each rule consists of an l-value, an operator, and an r-value. The l-value and r-value are each the opening half of an HTML tag, potentially including attributes. The operator is either `=` or `+=`.

Nodes in the second HTML document that match the r-value are copied into the first HTML document as children of nodes that match the l-value. The `=` operator replaces all the l-value's children. The `+=` operator appends to the l-value's children, which allows more complex results when passing more than two arguments to `mergician`.

If no rules are given on the command-line, these defaults are used:

```
<article class="body"> = <body>
<div class="body"> = <body>
<section class="body"> = <body>
```

Merge algorithm
---------------

Mergician merges the second HTML document it's given into the first, then the third into the merged first and second, and so on until all the documents are processed.

For each pair of documents it merges, it follows these steps:

1. Merge `<title>` tags by appending " / " and the second `<title>` tag's text to the first `<title>` tag's text. (TODO: If there is no first `<title>` tag, copy the second `<title>` tag in unchanged. Also TODO: extend the [rules](#rules) to support customizing how the title text nodes are merged.)
2. Merge `<head>` tags by appending all the unique children from the second `<head>` tag to the first.
3. Iterate over the [rules](#rules) and, for each l-value matched in the first document and r-value matched in the second, replace (with `=`) or append to (with `+=`) the match in the first document with the content from the match in the second.

Markdown
--------

Lots of documents will be written first in Markdown, rendered to HTML, and then merged with other HTML. Mergician supports this directly: If it's given a file with the extension `.md`, it will render the Markdown before continuing with the merge algorithm.

In order to ensure subsequent edits to Markdown files don't overwrite edits made directly to the HTML files that were rendered from those Markdown files, Mergician processes Markdown files as follows:

1. Replace `.md` with `.html` at the end of the document's name.
2. If the HTML document (1) exists, hash its contents.
3. If a file with the name of the HTML document (1) prefixed with a `.` and suffixed with `.sha256` exists, read its contents.
4. If hashes (2) and (3) both exist and do not match, exit with an error to preserve the edited HTML.
5. If hash (2) exists but hash (3) doesn't, warn and continue.
6. Render the Markdown to HTML.
7. Redo hash (2) and write it to hash (3).
8. Then merge the HTML with other HTML, etc.

For example, taking `layout.html` from above and `article.md`:

```
Hello, world!
=============

Hello, world!
```

We can again use Mergician to merge them together:

```sh
mergician layout.html article.md
```

You get your article wrapped in your layout:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<link href="layout.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>Hello, world! — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/> 
<div class="body">
<article>
<h1>Hello, world!</h1>
<p>Hello, world!</p>
</article>
</div>
<br/>
<footer><p>© 2025</p></footer>
</body>
</html>
```

See also
--------

Mergician powers a whole suite of tools that manipulate HTML documents:

* [Deadlinks](https://github.com/rcrowley/deadlinks): Scan a document root directory for dead links
* [Electrostatic](https://github.com/rcrowley/electrostatic): Mergician-powered, pure-HTML CMS
* [Feed](https://github.com/rcrowley/feed): Scan a document root directory to construct an Atom feed
* [Frag](https://github.com/rcrowley/frag): Extract fragments of HTML documents
* [Sitesearch](https://github.com/rcrowley/sitesearch): Index a document root directory and serve queries to it in AWS Lambda
