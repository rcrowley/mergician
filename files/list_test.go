package files

import (
	"iter"
	"slices"
	"testing"
)

func TestListDuplicate(t *testing.T) {
	l := &List{}
	l.Add("test.html")
	l.Add("test.html")
	if paths := seqSlice(l.IterRelative()); !slices.Equal(paths, []string{"test.html"}) {
		t.Fatal(paths)
	}
}

func TestListHTMLThenMarkdown(t *testing.T) {
	l := &List{}
	l.Add("test.html")
	l.Add("test.md")
	if paths := seqSlice(l.IterRelative()); !slices.Equal(paths, []string{"test.md"}) {
		t.Fatal(paths)
	}
}

func TestListMarkdownThenHTML(t *testing.T) {
	l := &List{}
	l.Add("test.md")
	l.Add("test.html")
	if paths := seqSlice(l.IterRelative()); !slices.Equal(paths, []string{"test.md"}) {
		t.Fatal(paths)
	}
}

func TestListNotMarkdownOrHTML(t *testing.T) {
	l := &List{}
	l.Add("test.go")
	if paths := seqSlice(l.IterRelative()); !slices.Equal(paths, []string{}) {
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
	if paths := seqSlice(l.IterRelative()); !slices.Equal(paths, []string{"a.html", "b.html", "c.html", "d.html"}) {
		t.Fatal(paths)
	}
}

func seqSlice[T any](seq iter.Seq[T]) []T {
	var slice []T
	for t := range seq {
		slice = append(slice, t)
	}
	return slice
}
