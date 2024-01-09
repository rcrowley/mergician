package main

import (
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
	t.Log(html.String(out)) // TODO assertions
}

func TestParseMergeMarkdown(t *testing.T) {
	pathnames := []string{"html/template.html", "markdown/test.md"}
	in, err := parse(pathnames)
	if err != nil {
		t.Fatal(err)
	}
	out, err := html.Merge(in, html.DefaultRules())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(html.String(out)) // TODO assertions
}
