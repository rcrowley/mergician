package files

import (
	"fmt"
	"iter"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/rcrowley/mergician/html"
)

type List struct {
	mu    sync.Mutex
	paths []string
	root  string
}

func NewList(root string) *List {
	return &List{root: root}
}

func (l *List) Add(path string) {
	ext := filepath.Ext(path)
	if ext != ".htm" && ext != ".html" && ext != ".md" && ext != ".zip" {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	// If the Markdown variant of this path is in the list already,
	// we're done.
	mdPath := fmt.Sprint(strings.TrimSuffix(path, ext), ".md")
	i := sort.SearchStrings(l.paths, mdPath)
	if i < len(l.paths) && l.paths[i] == mdPath {
		return
	}

	// If the HTML variant of this path is in the list already, convert
	// its extension to this extension.
	htmlPath := fmt.Sprint(strings.TrimSuffix(path, ext), ".html")
	i = sort.SearchStrings(l.paths, htmlPath)
	if i < len(l.paths) && l.paths[i] == htmlPath {
		l.paths[i] = path
		return
	}

	i = sort.SearchStrings(l.paths, path)
	if i == len(l.paths) || l.paths[i] != path {
		l.paths = append(l.paths, "")
		copy(l.paths[i+1:], l.paths[i:])
		l.paths[i] = path
	}
}

func (l *List) IterQualified() iter.Seq[string] {
	root := l.root // copy on purpose
	return func(yield func(string) bool) {
		for path := range l.IterRelative() {
			if !yield(filepath.Join(root, path)) {
				return
			}
		}
	}
}

func (l *List) IterRelative() iter.Seq[string] {
	l.mu.Lock()
	defer l.mu.Unlock()
	paths := append(([]string)(nil), l.paths...) // copy on purpose
	return func(yield func(string) bool) {
		for _, path := range paths {
			if !yield(path) {
				return
			}
		}
	}
}

func (l *List) Parse() ([]*html.Node, error) {
	return ParseSlice(l.Paths())
}

func (l *List) Paths() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append(([]string)(nil), l.paths...)
}
