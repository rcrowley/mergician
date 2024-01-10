package markdown

import "testing"

func TestParseFile(t *testing.T) {
	d, err := ParseFile("test.md")
	if err != nil {
		t.Fatal(err)
	}
	if String(d) != "<h1>Hello, world!</h1>\n<p>Lovely day for a test, isn't it?</p>\n" {
		t.Fatalf("%#v", String(d))
	}
}