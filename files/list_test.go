package files

import (
	"slices"
	"testing"
)

func TestListDuplicate(t *testing.T) {
	l := &List{}
	l.Add("test.html")
	l.Add("test.html")
	if paths := l.RelativePaths(); !slices.Equal(paths, []string{"test.html"}) {
		t.Fatal(paths)
	}
}

func TestListHTMLThenMarkdown(t *testing.T) {
	l := &List{}
	l.Add("test.html")
	l.Add("test.md")
	if paths := l.RelativePaths(); !slices.Equal(paths, []string{"test.md"}) {
		t.Fatal(paths)
	}
}

func TestListMarkdownThenHTML(t *testing.T) {
	l := &List{}
	l.Add("test.md")
	l.Add("test.html")
	if paths := l.RelativePaths(); !slices.Equal(paths, []string{"test.md"}) {
		t.Fatal(paths)
	}
}

func TestListNotMarkdownOrHTML(t *testing.T) {
	l := &List{}
	l.Add("test.go")
	if paths := l.RelativePaths(); !slices.Equal(paths, []string{}) {
		t.Fatal(paths)
	}
}

func TestListParse(t *testing.T) {
	l := &List{}
	l.Add("html.html")
	l.Add("md.md")
	nodes, err := l.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 2 {
		t.Fatal(nodes)
	}
}

func TestListSorted(t *testing.T) {
	l := &List{}
	l.Add("b.html")
	l.Add("a.html")
	l.Add("d.html")
	l.Add("c.html")
	if paths := l.RelativePaths(); !slices.Equal(paths, []string{"a.html", "b.html", "c.html", "d.html"}) {
		t.Fatal(paths)
	}
}

func TestListRoot(t *testing.T) {
	l := NewList("root")
	if root := l.Root(); root != "root" {
		t.Fatal(root)
	}
}
