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

// ParseFiles accepts a slice of HTML (or Markdown) filenames, parses the
// documents, and returns []*Node and a nil error or nil and a non-nil error.
func ParseFiles(pathnames []string) ([]*Node, error) {
	in := make([]*Node, len(pathnames))
	for i, pathname := range pathnames {
		var err error
		if in[i], err = ParseFile(pathname); err != nil {
			return nil, err
		}
	}
	return in, nil
}

// ParseString parses an HTML document or document fragment from a string.
func ParseString(s string) (*Node, error) {

	// If this looks like a complete document, treat it like one.
	if strings.HasPrefix(s, "<!") {
		return Parse(strings.NewReader(s))
	}

	// Otherwise, treat it as a fragment and try to find it in the document
	// the parser builds around it.
	nodes, err := html.ParseFragment(strings.NewReader(s), nil)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, fmt.Errorf("parse error: %s", s)
	}
	if len(nodes) > 1 {
		return nil, fmt.Errorf("fragment not rooted at one node: %s", s)
	}
	n := nodes[0]

	// If we're literally looking for <head>, return it.
	ok, err := regexp.MatchString("^<head[> ]", s)
	if err != nil {
		panic(err)
	}
	if ok {
		return n.FirstChild, nil
	}

	// If we're literally looking for <body>, return it.
	ok, err = regexp.MatchString("^<body[> ]", s)
	if err != nil {
		panic(err)
	}
	if ok {
		return n.FirstChild.NextSibling, nil
	}

	// Otherwise, usually, we're looking for whatever node has been added to
	// either the <head> or the <body> of the document the parser made up.
	if n.FirstChild.FirstChild != nil {
		return n.FirstChild.FirstChild, nil
	} else if n.FirstChild.NextSibling.FirstChild != nil {
		return n.FirstChild.NextSibling.FirstChild, nil
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
