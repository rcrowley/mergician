package html

import (
	"fmt"

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

	// <title> is special and I'm not sure what syntax to offer to expose it.
	if IsAtom(atom.Title)(dst) {
		Debug("processing <title>")
		if srcTitle := Find(src, IsAtom(atom.Title)); srcTitle != nil {
			for n := srcTitle.FirstChild; n != nil; n = n.NextSibling {
				dst.InsertBefore(NewTextNode(" \u2014 "), dst.FirstChild) // " &mdash; " but literal
				dst.InsertBefore(CopyNode(n), dst.FirstChild)
			}
		}
	}

	// <head> += <head>, except for <title>, which is ignored here, and
	// duplicate <link> and <meta> tags, which are deduplicated.
	if IsAtom(atom.Head)(dst) {
		Debug("processing <head>")
		var dedupe []string
		isLinkOrMeta := IsAtom(atom.Link, atom.Meta) // tags to dedupe
		for n := dst.FirstChild; n != nil; n = n.NextSibling {
			if isLinkOrMeta(n) {
				dedupe, _ = InsertSorted(dedupe, String(n))
			}
		}
		if srcHead := Find(src, IsAtom(atom.Head)); srcHead != nil {
			for n := srcHead.FirstChild; n != nil; n = n.NextSibling {
				if n == srcHead.FirstChild && IsWhitespace(n) {
					continue // assume we already have a "\n" within <head> from dst
				}
				var skipNextSiblingWhitespaceNode bool

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
				} else {
					dst.AppendChild(CopyNode(n))
				}

				if skipNextSiblingWhitespaceNode && n.NextSibling != nil && IsWhitespace(n.NextSibling) {
					n = n.NextSibling
					Debugf("skipping next sibling whitespace node: %#v", n.Data)
				}
			}
		}
	}

	// If this destination node matches the l-value in one of our rules, look
	// for a source node that matches its corresponding r-value and copy
	// it into place.
	for _, rule := range rules {
		//Debugf("processing %v", rule)
		if Match(rule.Dst)(dst) {
			if srcBody := Find(src, Match(rule.Src)); srcBody != nil {

				switch rule.Op {
				case "=":
					dst.FirstChild, dst.LastChild = nil, nil
				case "+=":
					if dst.LastChild != nil && IsWhitespace(dst.LastChild) {
						Debugf("removing last child whitespace node: %#v", dst.LastChild.Data)
						dst.RemoveChild(dst.LastChild)
					}
				default:
					panic(fmt.Sprintf(`html.Merge: operator in %v is not "=" or "+="`, rule))
				}

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
