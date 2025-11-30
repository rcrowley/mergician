package html

import (
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestParseFile(t *testing.T) {
	n, err := ParseFile("testdata/template.html")
	if err != nil {
		t.Fatal(err)
	}
	//printNodeAsTree(n, "")
	if n.Type != html.DocumentNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}

func TestParseFiles(t *testing.T) {
	nodes, err := ParseFiles([]string{"testdata/article.html", "testdata/template.html"})
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 2 {
		t.Fatal(nodes)
	}
	for _, n := range nodes {
		//printNodeAsTree(n, "")
		if n.Type != html.DocumentNode {
			t.Fatal(NodeTypeString(n.Type))
		}
	}
}

func TestParseStringDocument(t *testing.T) {
	n, err := ParseString(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>TestParseStringDocument</title>
</head>
<body>
<h1>TestParseStringDocument</h1>
</body>
</html>
`)
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(n, "")
	if n.Type != DocumentNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}

func TestParseStringDocumentPreWithBlankLines(t *testing.T) {
	n, err := ParseString(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>TestParseStringDocumentPreWithBlankLines</title>
</head>
<body>
<pre>


</pre>
</body>
</html>
`)
	if err != nil {
		t.Fatal(err)
	}
	n = Find(n, IsAtom(atom.Pre))
	if n == nil {
		t.Fatal(n)
	}

	actual := String(n)
	expected := "<pre>\n\n\n</pre>"
	if actual != expected {
		t.Fatalf("actual: %#v != expected: %#v", actual, expected)
	}
}

func TestParseStringFragmentBody(t *testing.T) {
	n, err := ParseString(`<body>`)
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(n, "")
	if n.DataAtom != atom.Body {
		t.Fatal(n.DataAtom, n.Data)
	}
	if n.Type != ElementNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}

func TestParseStringFragmentH1(t *testing.T) {
	n, err := ParseString(`<h1>TestParseStringFragment</h1>`)
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(n, "")
	if n.DataAtom != atom.H1 {
		t.Fatal(n.DataAtom, n.Data)
	}
	if n.Type != ElementNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}

func TestParseStringFragmentHead(t *testing.T) {
	n, err := ParseString(`<head>`)
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(n, "")
	if n.DataAtom != atom.Head {
		t.Fatal(n.DataAtom, n.Data)
	}
	if n.Type != ElementNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}

func TestParseStringFragmentMeta(t *testing.T) {
	n, err := ParseString(`<meta charset="utf-8">`)
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(n, "")
	if n.DataAtom != atom.Meta {
		t.Fatal(n.DataAtom, n.Data)
	}
	if n.Type != ElementNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}
