package files

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	l, err := All([]string{"."}, []string{}, []string{".html", ".md", ".zip"}) // skipping ".htm"
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.md", "test/html.html", "zip.zip"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTML(t *testing.T) {
	l, err := AllHTML([]string{"."}, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html", "test/html.html", "test/test/html.htm"}) {
		t.Fatal(pathnames)
	}
}

func TestAllHTMLExclude(t *testing.T) {
	l, err := AllHTML([]string{"."}, []string{"test"})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.html"}) {
		t.Fatal(pathnames)
	}
}

func TestAllInputs(t *testing.T) {
	l, err := AllInputs([]string{"."}, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if pathnames := l.Pathnames(); !slices.Equal(pathnames, []string{"html.html", "md.md", "test/html.html", "test/test/html.htm", "zip.zip"}) {
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
