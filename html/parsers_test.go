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
	//printNodeAsTree(n, "")
	if n.Type != html.DocumentNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}

func TestParseStringDocument(t *testing.T) {
	n, err := ParseString(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>TestPareStringDocument</title>
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

func TestParseStringFragment(t *testing.T) {
	n, err := ParseString(`<h1>TestParseStringFragment</h1>`)
	if err != nil {
		t.Fatal(err)
	}
	printNodeAsTree(n, "")
	if n.Type != ElementNode {
		t.Fatal(NodeTypeString(n.Type))
	}
}
