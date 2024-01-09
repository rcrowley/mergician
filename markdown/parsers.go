package markdown

import (
	"io"
	"os"
)

// Parse reads a complete Markdown document from an io.Reader. It is the
// caller's responsibility to ensure the io.Reader is positioned at the
// beginning of the document and to clean up (i.e. close file descriptors,
// etc.) afterwards. Most callers will want to use ParseFile instead.
func Parse(r io.Reader) (*Document, error) {
	source, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	d := &Document{}
	if err := markdown().Convert(source, d); err != nil {
		return nil, err
	}
	return d, nil
}

// ParseFile opens a Markdown file, parses the document it contains, closes
// the file descriptor, and returns the parsed Markdown document. In case of
// error, a nil *Node is returned along with the error.
func ParseFile(pathname string) (*Document, error) {
	f, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}
