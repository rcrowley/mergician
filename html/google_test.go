package html

import "testing"

func TestGoogle(t *testing.T) {
	n, err := Google("google.zip")
	if err != nil {
		t.Fatal(err)
	}
	//printNodeAsTree(n, "")

	actual := String(n)
	expected := `<html><head></head><body><h1>Heading 1</h1><p>Paragraph.</p><h2>Heading 2</h2><p>Paragraph with a <a href="https://example.com">link</a> and <strong>bold</strong>, <em>italic</em>, and <strong><em>bolditalic</em></strong> text.</p><ul><li>Also an</li><li>Unordered list</li></ul><p>Paragraph.</p><ol start="1"><li>And an</li><li>Ordered list</li></ol></body></html>`
	if actual != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}
}
