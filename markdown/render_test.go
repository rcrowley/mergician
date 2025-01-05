package markdown

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

func TestRenderFile(t *testing.T) {
	var (
		d      *Document
		err    error
		actual []byte
	)
	expected := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>Hello, world!</title>
</head>
<body>
<article>
<h1>Hello, world!</h1>
<p>Lovely day for a test, isn&rsquo;t it?</p>
</article>
</body>
</html>
`

	// No HTML or hash files in the way.
	if err := os.Remove("testdata/test.html"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if err := os.Remove("testdata/.test.html.sha256"); err != nil && !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	if d, err = ParseFile("testdata/test.md"); err != nil {
		t.Fatal(err)
	}
	if err := RenderFile("testdata/test.html", d); err != nil {
		t.Fatal(err)
	}
	if actual, err = os.ReadFile("testdata/test.html"); err != nil {
		t.Fatal(err)
	}
	if string(actual) != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	// Unmodified HTML in the way, which should still succeed.
	if d, err = ParseFile("testdata/test.md"); err != nil {
		t.Fatal(err)
	}
	if err := RenderFile("testdata/test.html", d); err != nil {
		t.Fatal(err)
	}
	if actual, err = os.ReadFile("testdata/test.html"); err != nil {
		t.Fatal(err)
	}
	if string(actual) != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

	// Modified HTML in the way, which should error.
	if err := os.Truncate("testdata/test.html", 0); err != nil {
		t.Fatal(err)
	}
	if err := RenderFile("testdata/test.html", d); err == nil {
		t.Fatal("expected error")
	}
	fi, err := os.Stat("testdata/test.html")
	if err != nil {
		t.Fatal(err)
	}
	if fi.Size() != 0 {
		t.Fatal("expected test.html to be empty")
	}

}
