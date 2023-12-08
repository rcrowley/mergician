package html

import (
	"io"
	"os"

	"golang.org/x/net/html"
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
	f, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}
