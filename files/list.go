package files

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/rcrowley/mergician/html"
)

type List struct {
	mu        sync.Mutex
	pathnames []string
}

func (l *List) Add(pathname string) {
	ext := filepath.Ext(pathname)
	if ext != ".html" && ext != ".md" {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	// If the Markdown variant of this pathname is in the list already,
	// we're done.
	mdPathname := fmt.Sprint(strings.TrimSuffix(pathname, ext), ".md")
	//log.Printf("Markdown %v", mdPathname)
	i := sort.SearchStrings(l.pathnames, mdPathname)
	if i < len(l.pathnames) && l.pathnames[i] == mdPathname {
		//log.Print("already got one")
		return
	}

	// If the HTML variant of this pathname is in the list already, convert
	// its extension to this extension.
	htmlPathname := fmt.Sprint(strings.TrimSuffix(pathname, ext), ".html")
	//log.Printf("HTML %v", htmlPathname)
	i = sort.SearchStrings(l.pathnames, htmlPathname)
	if i < len(l.pathnames) && l.pathnames[i] == htmlPathname {
		//log.Printf("replacing %v with %v", htmlPathname, pathname)
		l.pathnames[i] = pathname
		return
	}

	i = sort.SearchStrings(l.pathnames, pathname)
	if i == len(l.pathnames) || l.pathnames[i] != pathname {
		l.pathnames = append(l.pathnames, "")
		copy(l.pathnames[i+1:], l.pathnames[i:])
		l.pathnames[i] = pathname
	}
}

func (l *List) Parse() ([]*html.Node, error) {
	return ParseSlice(l.Pathnames())
}

func (l *List) Pathnames() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append(([]string)(nil), l.pathnames...)
}
