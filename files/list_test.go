package files

import "testing"

func TestListDuplicate(t *testing.T) {
	l := &List{}
	l.Add("test.html")
	l.Add("test.html")
	pathnames := l.Pathnames()
	if len(pathnames) != 1 || pathnames[0] != "test.html" {
		t.Fatal(pathnames)
	}
}

func TestListHTMLThenMarkdown(t *testing.T) {
	l := &List{}
	l.Add("test.html")
	l.Add("test.md")
	pathnames := l.Pathnames()
	if len(pathnames) != 1 || pathnames[0] != "test.md" {
		t.Fatal(pathnames)
	}
}

func TestListMarkdownThenHTML(t *testing.T) {
	l := &List{}
	l.Add("test.md")
	l.Add("test.html")
	pathnames := l.Pathnames()
	if len(pathnames) != 1 || pathnames[0] != "test.md" {
		t.Fatal(pathnames)
	}
}

func TestListNotMarkdownOrHTML(t *testing.T) {
	l := &List{}
	l.Add("test.go")
	pathnames := l.Pathnames()
	if len(pathnames) != 0 {
		t.Fatal(pathnames)
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
	pathnames := l.Pathnames()
	if len(pathnames) != 4 || pathnames[0] != "a.html" || pathnames[1] != "b.html" || pathnames[2] != "c.html" || pathnames[3] != "d.html" {
		t.Fatal(pathnames)
	}
}
