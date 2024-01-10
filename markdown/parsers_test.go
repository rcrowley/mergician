package markdown

import "testing"

func TestParseFile(t *testing.T) {
	d, err := ParseFile("test.md")
	if err != nil {
		t.Fatal(err)
	}
	if d.String() != "<h1>Hello, world!</h1>\n<p>Lovely day for a test, isn&rsquo;t it?</p>\n" {
		t.Fatalf("%#v", d.String()) // d.String() only works because type Document = bytes.Buffer
	}
}
