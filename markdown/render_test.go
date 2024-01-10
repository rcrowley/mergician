package markdown

import (
	"os"
	"testing"
)

func TestRenderFile(t *testing.T) {
	d, err := ParseFile("test.md")
	if err != nil {
		t.Fatal(err)
	}

	// No HTML or hash files in the way.
	if err := os.Remove("test.html"); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(".test.html.sha256"); err != nil {
		t.Fatal(err)
	}
	if err := RenderFile("test.html", d); err != nil {
		t.Fatal(err)
	}

	// Unmodified HTML in the way, which should still succeed.
	if err := RenderFile("test.html", d); err != nil {
		t.Fatal(err)
	}

	// Modified HTML in the way, which should error.
	if err := os.Truncate("test.html", 0); err != nil {
		t.Fatal(err)
	}
	if err := RenderFile("test.html", d); err == nil {
		t.Fatal("expected error")
	}

}
