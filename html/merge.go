package html

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Merge several parsed HTML documents into one. The second node is merged
// into a copy of the first, then the third into the copy, and so on until all
// arguments are processed. If any step results in an error, processing stops
// and a nil *Node is returned along with the error.
func Merge(in []*Node, rules []Rule) (*Node, error) {
	if len(in) == 0 {
		panic("html.Merge called with zero inputs")
	}
	Debug("initializing the output HTML document")
	out := CopyNode(in[0])
	if len(in) == 1 {
		return out, nil
	}

	for i := 1; i < len(in); i++ {
		Debugf("merging input HTML document %d into the output HTML document", i)
		if err := merge(out, in[i], rules); err != nil {
			return nil, err
		}
	}
	return out, nil
}

type MergeError string

func sprintfMergeError(format string, args ...interface{}) error {
	return MergeError(fmt.Sprintf(format, args...))
}

func (err MergeError) Error() string { return string(err) }

func merge(dst, src *Node, rules []Rule) error {
	//Debugf("merging dst: %s %q, src: %s %q", NodeTypeString(dst.Type), dst.Data, NodeTypeString(src.Type), src.Data)

	// Add a "\n" between <!DOCTYPE html> and <html>, after <html>, and after
	// </body>, none of which for some reason are included by the parser
	// despite definitely being present in the source files.
	if dst.Type == html.ElementNode && dst.DataAtom == atom.Html {
		dst.Parent.InsertBefore(NewTextNode("\n"), dst)     // between <!DOCTYPE html> and <html>
		dst.InsertBefore(NewTextNode("\n"), dst.FirstChild) // after <html>
		dst.AppendChild(NewTextNode("\n"))                  // after </body>
	}

	// <title> is special and I'm not sure what syntax to offer to expose it.
	if IsAtom(atom.Title)(dst) {
		if srcTitle := Find(src, IsAtom(atom.Title)); srcTitle != nil {
			for n := srcTitle.FirstChild; n != nil; n = n.NextSibling {
				dst.InsertBefore(&Node{
					Data: " \u2014 ", // " &mdash; " but literal
					Type: html.TextNode,
				}, dst.FirstChild)
				dst.InsertBefore(CopyNode(n), dst.FirstChild)
			}
		}
	}

	// <head> += <head>, except for <title>, which is ignored here, and
	// duplicate <link> and <meta> tags, which are deduplicated.
	if IsAtom(atom.Head)(dst) {
		var dedupe []string
		isLinkOrMeta := IsAtom(atom.Link, atom.Meta) // tags to dedupe
		for n := dst.FirstChild; n != nil; n = n.NextSibling {
			if isLinkOrMeta(n) {
				dedupe, _ = InsertSorted(dedupe, String(n))
			}
		}
		if srcHead := Find(src, IsAtom(atom.Head)); srcHead != nil {
			var skipNextSiblingWhitespaceNode bool
			for n := srcHead.FirstChild; n != nil; n = n.NextSibling {
				if skipNextSiblingWhitespaceNode {
					skipNextSiblingWhitespaceNode = false
					if IsWhitespace(n) {
						Debugf("skipping next sibling whitespace node: %#v", n.Data)
						continue
					}
				}
				if isLinkOrMeta(n) {
					var ok bool
					if dedupe, ok = InsertSorted(dedupe, String(n)); ok {
						dst.AppendChild(CopyNode(n))
					} else {
						Debugf("skipping duplicate <link> or <meta> tag: %v", n)
						skipNextSiblingWhitespaceNode = true
					}
				} else if IsAtom(atom.Title)(n) { // handled above
					Debug("skipping already-handled <title> tag")
					skipNextSiblingWhitespaceNode = true
				} else if !IsWhitespace(n) {
					dst.AppendChild(CopyNode(n))
				}
			}
		}
	}

	// If this destination node matches the l-value in one of our rules, look
	// for a source node that matches its corresponding r-value and copy
	// it into place.
	for _, rule := range rules {
		if Match(rule.Dst)(dst) {
			if srcBody := Find(src, Match(rule.Src)); srcBody != nil {

				// TODO this is rule.Op == "="; need to support rule.Op == "+=", too
				// TODO which might be as easy as not doing this line
				dst.FirstChild, dst.LastChild = nil, nil

				for n := srcBody.FirstChild; n != nil; n = n.NextSibling {
					dst.AppendChild(CopyNode(n))
				}
			}
		}
	}

	for dstChild := dst.FirstChild; dstChild != nil; dstChild = dstChild.NextSibling {
		merge(dstChild, src, rules)
	}

	//Debugf("merged dst: %s %q, src: %s %q", NodeTypeString(dst.Type), dst.Data, NodeTypeString(src.Type), src.Data)
	return nil
}
