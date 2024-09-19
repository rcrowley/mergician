package main

import (
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/rcrowley/mergician/html"
)

func TestParseMergeHTML(t *testing.T) {
	pathnames := []string{"html/template.html", "html/article.html"}
	in, err := parse(pathnames)
	if err != nil {
		t.Fatal(err)
	}
	out, err := html.Merge(in, html.DefaultRules())
	if err != nil {
		t.Fatal(err)
	}

	actual := html.String(out)
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
<link href="template.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>My cool webpage — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/><!-- explicit self-closing -->
<article class="body">
<h1>Things</h1>
<p>Stuff</p>
</article>
<br/><!-- implied self-closing -->
<footer><p>© 2023</p></footer>
</body>
</html>
`
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(html.String(out))
}

func TestParseMergeMarkdown(t *testing.T) {

	// Remove the HTML and hash files. We're cheating and reusing a file from
	// another test and this ensures we don't encounter conflicts.
	if err := os.Remove("test.html"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if err := os.Remove(".test.html.sha256"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}

	pathnames := []string{"html/template.html", "test.md"}
	in, err := parse(pathnames)
	if err != nil {
		t.Fatal(err)
	}
	out, err := html.Merge(in, html.DefaultRules())
	if err != nil {
		t.Fatal(err)
	}

	actual := html.String(out)
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
<link href="template.css" rel="stylesheet"/>
<meta charset="utf-8"/>
<meta content="width=device-width,initial-scale=1" name="viewport"/>
<title>Hello, world! — Website</title>
</head>
<body>
<header><h1>Website</h1></header>
<br/><!-- explicit self-closing -->
<article class="body">
<article>
<h1>Hello, world!</h1>
<p>Lovely day for a test, isn’t it?</p>
</article>
</article>
<br/><!-- implied self-closing -->
<footer><p>© 2023</p></footer>
</body>
</html>
` // <article>-in-<article> is weird but just an artifact of this specific test harness
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(html.String(out))
}
