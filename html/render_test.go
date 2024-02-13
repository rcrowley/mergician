package html

import (
	"os"
	"testing"
)

func TestRenderFile(t *testing.T) {
	n, err := ParseFile("article.html")
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
	expected := `<!DOCTYPE html><html><head>
<link href="template.css" rel="stylesheet"/>
<title>My cool webpage</title>
</head>
<body>
<h1>Things</h1>
<p>Stuff</p>


</body></html>
` // TODO deal with the deranged whitespace the merge algorithm produces
	if string(actual) != expected {
		t.Fatalf("actual: %s != expected: %s", actual, expected)
	}

}
