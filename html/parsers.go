package html

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Parse reads a complete HTML document from an io.Reader. It is the caller's
// responsibility to ensure the io.Reader is positioned at the beginning of
// the document and to clean up (i.e. close file descriptors, etc.) afterwards.
// Most callers will want to use ParseFile instead.
func Parse(r io.Reader) (n *Node, err error) {
	if n, err = html.Parse(r); err != nil {
		return
	}

	// Be pendantic about the whitespace nodes that go into the parse tree so
	// that we can be more easily pedantic about the HTML this renders later.
	if n.Type == DocumentNode && n.FirstChild != nil && n.FirstChild.Type == DoctypeNode {
		htmlNode := n.FirstChild.NextSibling
		if htmlNode != nil && htmlNode.Type == ElementNode && htmlNode.DataAtom == atom.Html {
			n.InsertBefore(NewTextNode("\n"), htmlNode)                   // between <!DOCTYPE html> and <html>
			htmlNode.InsertBefore(NewTextNode("\n"), htmlNode.FirstChild) // between <html> and <head>
			htmlNode.AppendChild(NewTextNode("\n"))                       // between </body> and </html>
			n.AppendChild(NewTextNode("\n"))                              // after </html>
		}
	}
	removeConsecutiveNewlines(n)

	return
}

// ParseFile opens an HTML file, parses the document it contains, closes the
// file descriptor, and returns the parsed HTML document. In case of error,
// a nil *Node is returned along with the error.
func ParseFile(pathname string) (*Node, error) {

	// Detect and descend into Google Docs' HTML-in-zip format.
	// TODO also support EPUB files
	if ext := filepath.Ext(pathname); ext == ".zip" {
		return Google(pathname)
	}

	f, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}

// ParseString parses an HTML document fragment from a string. It is intended
// to be used on <body> elements or fragments that can be contained within a
// <body> element. It does not work for full documents contained in strings.
// For those cases, use Parse(strings.NewReader(s)).
func ParseString(s string) (*Node, error) {
	n, err := html.Parse(strings.NewReader(""))
	if err != nil {
		panic(err)
	}
	nodes, err := html.ParseFragment(strings.NewReader(s), n.FirstChild.FirstChild.NextSibling)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		nodes, err = html.ParseFragment(strings.NewReader(s), n.FirstChild)
		if len(nodes) == 0 {
			return nil, fmt.Errorf("parse error: %s", s)
		}
	}
	for _, n := range nodes {
		if IsAtom(atom.Head)(n) {
			continue
		}
		return n, nil
	}
	return nil, fmt.Errorf("parse error: %s", s)
}

var consecutiveNewlines = regexp.MustCompile("\n\n+")

func removeConsecutiveNewlines(n *Node) {
	if IsWhitespace(n) {
		n.Data = consecutiveNewlines.ReplaceAllString(n.Data, "\n")
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		removeConsecutiveNewlines(child)
	}
}
