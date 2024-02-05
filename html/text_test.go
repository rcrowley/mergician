package html

import (
	"encoding/json"
	"testing"
)

func TestText(t *testing.T) {
	n, err := ParseFile("article.html")
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

func jsonString(document any) string {
	b, err := json.MarshalIndent(document, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(b)
}
