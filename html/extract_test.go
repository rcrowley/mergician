package html

import "testing"

func TestExtractFirstParagraph(t *testing.T) {
	n, err := ParseFile("article.html")
	if err != nil {
		t.Fatal(err)
	}
	if p := FirstParagraph(n); p != "Stuff" {
		t.Fatalf("%#v", p)
	}
}

func TestExtractTitle(t *testing.T) {
	n, err := ParseFile("article.html")
	if err != nil {
		t.Fatal(err)
	}
	if title := Title(n); title != "My cool webpage" {
		t.Fatalf("%#v", title)
	}
}
