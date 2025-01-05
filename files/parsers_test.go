package files

import (
	"testing"

	"github.com/rcrowley/mergician/html"
)

func TestParseHTML(t *testing.T) {
	n, err := Parse("testdata/html.html")
	if err != nil {
		t.Fatal(err)
	}
	if n.Type != html.DocumentNode {
		t.Fatal(html.NodeTypeString(n.Type))
	}
}

func TestParseMarkdown(t *testing.T) {
	n, err := Parse("testdata/md.md")
	if err != nil {
		t.Fatal(err)
	}
	if n.Type != html.DocumentNode {
		t.Fatal(html.NodeTypeString(n.Type))
	}
}

func TestParseSlice(t *testing.T) {
	nodes, err := ParseSlice([]string{"testdata/html.html", "testdata/md.md"})
	if err != nil {
		t.Fatal(err)
	}
	if len(nodes) != 2 {
		t.Fatal(nodes)
	}
	for _, n := range nodes {
		if n.Type != html.DocumentNode {
			t.Fatal(html.NodeTypeString(n.Type))
		}
	}
}
