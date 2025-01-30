package html

import (
	"os"
	"testing"
)

func TestRenderFile(t *testing.T) {
	n, err := ParseFile("testdata/article.html")
	if err != nil {
		t.Fatal(err)
	}
	if err := RenderFile("test.html", n); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.html")

	actual, err := os.ReadFile("test.html")
	if err != nil {
		t.Fatal(err)
	}
	expected := testArticleRenderedHTML(t)
	if string(actual) != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}
}

func testArticleRenderedHTML(t *testing.T) string {
	b, err := os.ReadFile("testdata/article.rendered.html")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
