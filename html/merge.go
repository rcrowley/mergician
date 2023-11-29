package html

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// Merge several parsed HTML documents into one. The second argument is merged
// into a copy of the first, then the third into the copy, and so on until all
// arguments are processed. If any step results in an error, processing stops
// and a nil *Node is returned along with the error.
func Merge(in ...*Node) (*Node, error) {
	if len(in) == 0 {
		panic("html.Merge called with zero arguments")
	}
	out := copyNode(in[0])
	if len(in) == 1 {
		return out, nil
	}

	for i := 1; i < len(in); i++ {
		if err := merge(out, in[i]); err != nil {
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

func copyNode(in *Node) (out *Node) {
	out = &Node{
		Attr:      make([]html.Attribute, len(in.Attr)),
		DataAtom:  in.DataAtom,
		Data:      in.Data,
		Namespace: in.Namespace,
		Type:      in.Type,
	}
	for i := 0; i < len(in.Attr); i++ {
		out.Attr[i] = in.Attr[i]
	}
	for n := in.FirstChild; n != nil; n = n.NextSibling {
		out.AppendChild(copyNode(n))
	}
	return
}

func merge(dst, src *Node) error {

	// XXX EXPLAIN YOURSELF
	if IsAtom(atom.Title)(dst) {
		if srcTitle := Find(src, IsAtom(atom.Title)); srcTitle != nil {
			for n := srcTitle.FirstChild; n != nil; n = n.NextSibling {
				dst.AppendChild(&Node{Data: " / ", Type: html.TextNode})
				dst.AppendChild(copyNode(n))
			}
		}
	}
	if IsAtom(atom.Head)(dst) {
		if srcHead := Find(src, IsAtom(atom.Head)); srcHead != nil {
			for n := srcHead.FirstChild; n != nil; n = n.NextSibling {
				if !IsAtom(atom.Title)(n) {
					dst.AppendChild(copyNode(n))
				}
			}
		}
	}
	if IsAtom(atom.Article, atom.Div, atom.Section)(dst) && HasAttr("class", "body")(dst) {
		if srcBody := Find(src, IsAtom(atom.Body)); srcBody != nil {
			dst.FirstChild, dst.LastChild = nil, nil
			for n := srcBody.FirstChild; n != nil; n = n.NextSibling {
				dst.AppendChild(copyNode(n))
			}
		}
	}

	// TODO suppress merging a "\n" text node immediately after another "\n" text node

	for dstChild := dst.FirstChild; dstChild != nil; dstChild = dstChild.NextSibling {
		merge(dstChild, src)
	}
	return nil
}

func nodeTypeString(t html.NodeType) string {
	switch t {
	case html.CommentNode:
		return "CommentNode"
	case html.DoctypeNode:
		return "DoctypeNode"
	case html.DocumentNode:
		return "DocumentNode"
	case html.ElementNode:
		return "ElementNode"
	case html.ErrorNode:
		return "ErrorNode"
	case html.RawNode:
		return "RawNode"
	case html.TextNode:
		return "TextNode"
	default:
		return fmt.Sprint(t)
	}
}
