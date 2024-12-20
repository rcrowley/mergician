package files

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	l, err := All([]string{"."}, []string{}, []string{".html", ".md", ".zip"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.md", "test/html.html", "test/test/html.html", "zip.zip"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTML(t *testing.T) {
	l, err := All([]string{"."}, []string{}, []string{".html"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html", "test/html.html", "test/test/html.html"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTMLExclude(t *testing.T) {
	l, err := All([]string{"."}, []string{"test"}, []string{".html"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html"}) {
		t.Fatal(pathnames)
	}
}

func TestAllMarkdown(t *testing.T) {
	l, err := All([]string{"."}, []string{}, []string{".md"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"md.md"}) {
		t.Fatal(pathnames)
	}
}
