package files

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	lists, err := All([]string{"testdata"}, []string{}, []string{".html", ".md", ".zip"}) // skipping ".htm"
	if err != nil {
		t.Fatal(err)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"html.html", "md.md", "test/html.html", "test2/html.html", "zip.zip"}) { // no ".htm"
		t.Fatal(paths)
	}
}

func TestAllHTML(t *testing.T) {
	lists, err := AllHTML([]string{"testdata"}, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"html.html", "md.html", "test/html.html", "test/test/htm.htm", "test2/html.html"}) {
		t.Fatal(paths)
	}
}

func TestAllHTMLExclude(t *testing.T) {
	lists, err := AllHTML([]string{"testdata"}, []string{"testdata/test"})
	if err != nil {
		t.Fatal(err)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"html.html", "md.html", "test2/html.html"}) {
		t.Fatal(paths)
	}
}

func TestAllHTMLExclude2(t *testing.T) {
	lists, err := AllHTML([]string{"testdata"}, []string{"testdata/test", "testdata/test2"})
	if err != nil {
		t.Fatal(err)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"html.html", "md.html"}) {
		t.Fatal(paths)
	}
}

func TestAllInputs(t *testing.T) {
	lists, err := AllInputs([]string{"testdata"}, []string{})
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 1 {
		t.Fatal(lists)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{
		"html.html", "md.md", "test/html.html", "test/test/htm.htm", "test2/html.html", "zip.zip",
	}) {
		t.Fatal(paths)
	}
}

func TestAllMarkdown(t *testing.T) {
	lists, err := All([]string{"testdata"}, []string{}, []string{".md"})
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 1 {
		t.Fatal(lists)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"md.md"}) {
		t.Fatal(paths)
	}
}

func TestAllSub(t *testing.T) {
	lists, err := All([]string{"testdata/test"}, []string{}, []string{".htm", ".html"})
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 1 {
		t.Fatal(lists)
	}
	if paths := lists[0].QualifiedPaths(); !slices.Equal(paths, []string{"testdata/test/html.html", "testdata/test/test/htm.htm"}) {
		t.Fatal(paths)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"html.html", "test/htm.htm"}) {
		t.Fatal(paths)
	}
}

func TestAllSub2(t *testing.T) {
	lists, err := All([]string{"testdata/test", "testdata/test2"}, []string{}, []string{".htm", ".html"})
	if err != nil {
		t.Fatal(err)
	}
	if len(lists) != 2 {
		t.Fatal(lists)
	}
	if paths := lists[0].QualifiedPaths(); !slices.Equal(paths, []string{"testdata/test/html.html", "testdata/test/test/htm.htm"}) {
		t.Fatal(paths)
	}
	if paths := lists[1].QualifiedPaths(); !slices.Equal(paths, []string{"testdata/test2/html.html"}) {
		t.Fatal(paths)
	}
	if paths := lists[0].RelativePaths(); !slices.Equal(paths, []string{"html.html", "test/htm.htm"}) {
		t.Fatal(paths)
	}
	if paths := lists[1].RelativePaths(); !slices.Equal(paths, []string{"html.html"}) {
		t.Fatal(paths)
	}
}
