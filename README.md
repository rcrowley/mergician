CMS that needs a better name
============================

I like writing plain HTML. I like it so much more than using all the other tools I've used for making websites that I've been known to use `sed`(1) to make bulk edits or over-compromise on navigation to avoid having to use a higher-level CMS. Anyway, this is my take on a truly minimalist CMS.

Consider these inputs:

`article.html`:

```html
<!DOCTYPE html>
<html>
<head><title>My cool webpage</title></head>
<body>
<h1>Things</h1>
<p>Stuff</p>
</body>
</html>
```

`template.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
</head>
<body>
<header><h1>Website</h1></header>
<footer><p>&copy; 2023</p></footer>
</body>
</html>
```

Then running `cms -srcdir . -dstdir public_html template.html article.html` (so ordered to make `xargs`(1)-like invocations more natural) will combine the two and write `public_html/article.html` to disk containing the following:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>My cool webpage</title>
</head>
<body>
<header><h1>Website</h1></header>
<h1>Things</h1>
<p>Stuff</p>
<footer><p>&copy; 2023</p></footer>
</body>
</html>
```

I still need to figure out how to indicate in the template's HTML how to order/interleave body elements, where order matters. I want the template to remain a viable HTML file itself but, more importantly, I really, really want to keep the inputs as viable HTML files.

I think this calls for microformats. Something like this:

`article.html`:

```html
<!DOCTYPE html>
<html>
<head><title>My cool webpage</title></head>
<body>
<article class="body">
<h1>Things</h1>
<p>Stuff</p>
</article>
</body>
</html>
```

`template.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
</head>
<body>
<header><h1>Website</h1></header>
<article class="body"></article>
<footer><p>&copy; 2023</p></footer>
</body>
</html>
```

And whatever the command is would merge those into this:

```html
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>My cool webpage</title>
</head>
<body>
<header><h1>Website</h1></header>
<article class="body">
<h1>Things</h1>
<p>Stuff</p>
</section>
<footer><p>&copy; 2023</p></footer>
</body>
</html>
```

This still doesn't feel quite right.

At any rate I'd probably build it with <https://pkg.go.dev/golang.org/x/net/html>.

----

`class="body"` should be something of a special case that pulls the _entire_ `<body>` of the other document into the tag bearing that class.

How would I build a feed or index?

* The order of the files given as input will matter a lot: `cms ... template.html first.html second.html third.html`
* That suggests we'd want a tool that looks inside a selection of files and sorts them by e.g. the first `<date>` tag it finds
* I wonder if we could do something with repetition of `class="body"` (since it's a class and can be repeated) to take each successive argument
* This wouldn't help use with input file lists of an unknown length, though maybe we could infer something by there being two `class="body"`

----

1. Process each document to construct an index of all the potential merge sites in each one. The index should map a merge site to an ordered list of `*Node`. Merge sites include:
    * `<head>`
    * `<body>`
    * `<article>`
    * `<article class="body">` (as an example of a container for a whole `<body>` from a document being merged in)
    * We might need a class prefix or something to distinguish merge sites from other uses of CSS
    * `<script>` in `<body>` (or will this be covered simply by `<body>`?)
    * et al
2. For each merge site, append all the children from input document, in order, into the output document. Do not append the merge site itself but rather each child. For example, don't append an input `<head>` to an output `<head>`, either inside or as a sibling; instead, append the e.g. `<link>`, `<meta>`, and `<title>` it contains into the output `<head>`.

----

Even if I don't implement it right out of the gate, I think there's a transmogrification language hiding in the informal spec I've been building to. The grammar would probably look something like this:

    Rules = Rule { Rules } .
    Rule = TagSpec [ "*" ] "->" TagSpec "\n" .
    TagSpec = "<" Name [ "class=\"" Name "\"" ] ">"
    Name = /^[a-z][a-z0-9]*$/ | "*"

The informal spec I've been sketching could be expressed as follows in this grammar:

    <head> * -> <head>
    <body> * -> <* class="body">

That might be it? I'm failing to remember why the `*` token is important.

----

No, that's insane. The syntax should be something like this encoding of what should probably be the default rules even before such a syntax for customization exists:

    <head> += <head> *
    <article class="body"> = <body>
    <section class="body"> = <body>
    <div class="body"> = <body>

Note that this new syntax is more of a reference to assignment in a typical programming language with typical l-values and r-values.
