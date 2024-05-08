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
	out := CopyNode(in[0])
	if len(in) == 1 {
		return out, nil
	}

	for i := 1; i < len(in); i++ {
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
	// duplicate <link> and <meta> tags.
	if IsAtom(atom.Head)(dst) {
		var dedupe []string
		for n := dst.FirstChild; n != nil; n = n.NextSibling {
			if IsAtom(atom.Link, atom.Meta)(n) { // tags to dedupe
				dedupe, _ = InsertSorted(dedupe, String(n))
			}
		}
		if srcHead := Find(src, IsAtom(atom.Head)); srcHead != nil {
			for n := srcHead.FirstChild; n != nil; n = n.NextSibling {
				if IsAtom(atom.Link, atom.Meta)(n) { // tags to dedupe
					var ok bool
					if dedupe, ok = InsertSorted(dedupe, String(n)); ok {
						dst.AppendChild(CopyNode(n))
					}
				} else if !IsAtom(atom.Title)(n) { // handled above
					dst.AppendChild(CopyNode(n))
				}
			}
		}
	}

	for _, rule := range rules {
		if Match(rule.Dst)(dst) {
			if srcBody := Find(src, Match(rule.Src)); srcBody != nil {
				// TODO this is rule.Op == "="; need to support rule.Op == "+=", too
				dst.FirstChild, dst.LastChild = nil, nil
				for n := srcBody.FirstChild; n != nil; n = n.NextSibling {
					dst.AppendChild(CopyNode(n))
				}
			}
		}
	}

	// TODO suppress merging a "\n" text node immediately after another "\n" text node

	for dstChild := dst.FirstChild; dstChild != nil; dstChild = dstChild.NextSibling {
		merge(dstChild, src, rules)
	}
	return nil
}
