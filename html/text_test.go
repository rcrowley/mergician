package html

import (
	"encoding/json"
	"testing"
)

func TestText(t *testing.T) {
	n, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	text := Text(n)
	if text.Nodes[0].Nodes[0].Nodes[0].Nodes[0].Text != "My cool webpage" {
		t.Fatal(jsonString(text))
	}
	if text.Nodes[0].Nodes[1].Nodes[0].Nodes[0].Text != "Things" {
		t.Fatal(jsonString(text))
	}
	if text.Nodes[0].Nodes[1].Nodes[1].Nodes[0].Text != "Stuff" {
		t.Fatal(jsonString(text))
	}
	//t.Log(jsonString(text))
}

func TestTextString(t *testing.T) {
	n, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	if s := Text(n).String(); s != "My cool webpage Things Stuff" {
		t.Fatalf("%#v", s)
	}
}

func TestTextZero(t *testing.T) {
	var text TextOnlyNode
	if s := text.String(); s != "" { // really the test is that this doesn't panic
		t.Fatalf("%#v", s)
	}
}

func jsonString(document any) string {
	b, err := json.MarshalIndent(document, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(b)
}
