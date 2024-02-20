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
	printNodeAsTree(n, "")
	if n.Type != html.DocumentNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}
