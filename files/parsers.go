package files

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/rcrowley/mergician/html"
	"github.com/rcrowley/mergician/markdown"
)

// Parse parses an HTML or Markdown file and returns a *html.Node and an error;
// only one of the two will be non-nil.
//
// If the pathname ends in ".md", it will parse the Markdown file and render
// the resulting HTML to a file in the same directory with the same name and
// the ".html" extension plus write the SHA256 sum of the HTML document to a
// file in the same directory with the same name and the ".html.sha256"
// extension (to detect changes).
func Parse(pathname string) (*html.Node, error) {
	if ext := filepath.Ext(pathname); ext == ".md" {
		d, err := markdown.ParseFile(pathname)
		if err != nil {
			return nil, err
		}
		pathname = fmt.Sprintf("%s.html", strings.TrimSuffix(pathname, ext))
		if err := markdown.RenderFile(pathname, d); err != nil {
			return nil, err
		}
	}
	return html.ParseFile(pathname)
}

// ParseSlice parses a slice of pathnames using Parse and returns a
// []*html.Node the same length as pathnames or nil and a non-nil error.
func ParseSlice(pathnames []string) ([]*html.Node, error) {
	nodes := make([]*html.Node, len(pathnames))
	for i, pathname := range pathnames {
		var err error
		if nodes[i], err = Parse(pathname); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}
