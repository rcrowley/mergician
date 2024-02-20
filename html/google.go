package html

import (
	"archive/zip"
	"fmt"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/net/html/atom"
)

// Google extracts a non-deranged HTML node from a zip file containing an HTML
// file obtained from the File > Download > Web Page (.html, zipped) option
// in Google Docs.
func Google(pathname string) (*Node, error) {
	n, err := parseZipFile(pathname)
	if err != nil {
		return nil, err
	}

	var strongClasses, emClasses []string
	if style := Find(n, IsAtom(atom.Style)); style != nil {
		strongClasses, emClasses = parseGoogleCSS(Text(style).String())
	}

	return filterGoogleNode(n, strongClasses, emClasses), nil
}

// filterGoogleNode removes the deranged parts of Google Docs' HTML and
// converts <span class="..."> into <strong> and <em> as the CSS dictates.
func filterGoogleNode(in *Node, strongClasses, emClasses []string) (out *Node) {

	// Step over <meta> tags. TODO only the content-type meta tag
	if in.DataAtom == atom.Meta {
		return nil
	}

	// Step over <style> tags because we don't want to look like paper.
	if in.DataAtom == atom.Style {
		return nil
	}

	out = &Node{
		DataAtom:  in.DataAtom,
		Data:      in.Data,
		Namespace: in.Namespace,
		Type:      in.Type,
	}
	for i, attr := range in.Attr {

		// Replace Google's wrapped URLs with the raw URLs.
		if in.DataAtom == atom.A && attr.Key == "href" {
			if u, err := url.Parse(attr.Val); err == nil {
				if v := u.Query().Get("q"); v != "" {
					attr.Val = v
				}
			}
		}

		// Omit class and id attributes since they're meaningless outside
		// of Google Docs. Use the class attribute to figure out whether
		// a <span class="..."> should be converted into <strong> or <em>.
		switch attr.Key {
		case "class":
			var strong, em bool
			for _, className := range strings.Split(attr.Val, " ") {
				if i := sort.SearchStrings(strongClasses, className); i < len(strongClasses) && strongClasses[i] == className {
					strong = true
				}
				if i := sort.SearchStrings(emClasses, className); i < len(emClasses) && emClasses[i] == className {
					em = true
				}
			}
			if strong && em {
				out.DataAtom = atom.Strong
				out.Data = "strong"
				in.Attr[i].Key = "mergician-masked-class"
				in.DataAtom = atom.Em
				in.Data = "em"
				if child := filterGoogleNode(in, strongClasses, emClasses); child != nil {
					out.AppendChild(child)
				}
				return
			} else if strong {
				out.DataAtom = atom.Strong
				out.Data = "strong"
			} else if em {
				out.DataAtom = atom.Em
				out.Data = "em"
			}
		case "id":
			// do nothing; these values are nonsense
		case "mergician-masked-class":
			// do nothing; this merely marks recursion termination
		default:
			out.Attr = append(out.Attr, attr)
		}

	}

	for n := in.FirstChild; n != nil; n = n.NextSibling {
		if child := filterGoogleNode(n, strongClasses, emClasses); child != nil {

			// Step into <span> tags but don't include the <span> tags themselves,
			// which are useless. (We've already switched important <span> tags to
			// more semantically meaningful tags.
			if child.DataAtom == atom.Span {
				for grandchild := child.FirstChild; grandchild != nil; grandchild = grandchild.NextSibling {
					grandchild.Parent = nil
					out.AppendChild(grandchild)
				}

			} else {
				out.AppendChild(child)
			}
		}
	}
	return
}

func parseGoogleCSS(raw string) (strongClasses, emClasses []string) {
	for _, directive := range strings.SplitAfter(raw, "}") {
		if strings.HasPrefix(directive, ".") {
			className := directive[1:strings.Index(directive, "{")]
			if strings.Contains(directive, "font-weight:700") {
				strongClasses = append(strongClasses, className)
			}
			if strings.Contains(directive, "font-style:italic") {
				emClasses = append(emClasses, className)
			}
		}
	}
	sort.Strings(strongClasses)
	sort.Strings(emClasses)
	return
}

func parseZipFile(pathname string) (*Node, error) {
	r, err := zip.OpenReader(pathname)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if filepath.Ext(f.Name) == ".html" {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()

			n, err := Parse(rc)
			if err != nil {
				return nil, err
			}
			return n, nil

		}
	}
	return nil, fmt.Errorf("no HTML file found in %s", pathname)
}
