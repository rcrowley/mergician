package main

import (
	"bytes"
	"errors"
	"io/fs"
	"os"
	"testing"

	"github.com/rcrowley/mergician/files"
	"github.com/rcrowley/mergician/html"
)

func TestMain(t *testing.T) {
	stdout := &bytes.Buffer{}
	Main([]string{"mergician", "html/testdata/template.html", "html/testdata/article.html"}, os.Stdin, stdout)
	actual := stdout.String()
	expected := testTemplatePlusArticleHTML(t)
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}
}

func TestParseMergeHTML(t *testing.T) {
	in, err := files.ParseSlice([]string{"html/testdata/template.html", "html/testdata/article.html"})
	if err != nil {
		t.Fatal(err)
	}
	out, err := html.Merge(in, html.DefaultRules())
	if err != nil {
		t.Fatal(err)
	}

	actual := html.String(out)
	expected := testTemplatePlusArticleHTML(t)
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(html.String(out))
}

func TestParseMergeMarkdown(t *testing.T) {

	// Remove the HTML and hash files. We're cheating and reusing a file from
	// another test and this ensures we don't encounter conflicts.
	if err := os.Remove("testdata/test.html"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if err := os.Remove("testdata/.test.html.sha256"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}

	in, err := files.ParseSlice([]string{"html/testdata/template.html", "testdata/test.md"})
	if err != nil {
		t.Fatal(err)
	}
	out, err := html.Merge(in, html.DefaultRules())
	if err != nil {
		t.Fatal(err)
	}

	actual := html.String(out)
	expected := testTemplatePlusTestHTML(t) // <article>-in-<article> is weird but just an artifact of this specific test harness
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	//t.Log(html.String(out))
}

func testTemplatePlusArticleHTML(t *testing.T) string {
	b, err := os.ReadFile("html/testdata/template+article.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func testTemplatePlusTestHTML(t *testing.T) string {
	b, err := os.ReadFile("testdata/template+test.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
