package html

import (
	"testing"

	"golang.org/x/net/html"
)

func TestParseFile(t *testing.T) {
	n, err := ParseFile("template.html")
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(t, n, "")
}

func printNodeAsTree(t *testing.T, n *Node, indent string) {
	switch n.Type {
	case html.ElementNode:
		t.Logf("%s<%s>\n", indent, n.DataAtom) // TODO n.Attr
	case html.TextNode:
		t.Logf("%s%q\n", indent, n.Data) // TODO n.Attr
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		printNodeAsTree(t, child, indent+"\t")
	}
	switch n.Type {
	case html.ElementNode:
		t.Logf("%s</%s>\n", indent, n.DataAtom)
	}
}
