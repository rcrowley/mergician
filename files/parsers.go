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
// If the path ends in ".md", it will parse the Markdown file and render the
// resulting HTML to a file in the same directory with the same name and the
// ".html" extension plus write the SHA256 sum of the HTML document to a file
// in the same directory with the same name and the ".html.sha256" extension
// (to detect changes).
func Parse(path string) (*html.Node, error) {
	if ext := filepath.Ext(path); ext == ".md" {
		d, err := markdown.ParseFile(path)
		if err != nil {
			return nil, err
		}
		path = fmt.Sprintf("%s.html", strings.TrimSuffix(path, ext))
		if err := markdown.RenderFile(path, d); err != nil {
			return nil, err
		}
	}
	return html.ParseFile(path)
}

// ParseSlice parses a slice of paths using Parse and returns a
// []*html.Node the same length as paths or nil and a non-nil error.
func ParseSlice(paths []string) ([]*html.Node, error) {
	nodes := make([]*html.Node, len(paths))
	for i, path := range paths {
		var err error
		if nodes[i], err = Parse(path); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}
