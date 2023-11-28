package html

import "testing"

func TestMergeSimple(t *testing.T) {
	n0, err := ParseFile("template.html")
	if err != nil {
		t.Fatal(err)
	}
	n1, err := ParseFile("article.html")
	if err != nil {
		t.Fatal(err)
	}
	n, err := Merge(n0, n1)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(String(n))
}
