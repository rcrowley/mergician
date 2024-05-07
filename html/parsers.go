package html

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Parse reads a complete HTML document from an io.Reader. It is the caller's
// responsibility to ensure the io.Reader is positioned at the beginning of
// the document and to clean up (i.e. close file descriptors, etc.) afterwards.
// Most callers will want to use ParseFile instead.
func Parse(r io.Reader) (*Node, error) {
	return html.Parse(r)
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
// <body> element.
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
