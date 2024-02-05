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
	t.Logf("%#v", html.String(out)) // TODO assertions
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
	t.Logf("%#v", html.String(out)) // TODO assertions
}
