package html

import (
	"fmt"
	"regexp"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	CommentNode  = html.CommentNode
	DoctypeNode  = html.DoctypeNode
	DocumentNode = html.DocumentNode
	ElementNode  = html.ElementNode
	ErrorNode    = html.ErrorNode
	RawNode      = html.RawNode
	TextNode     = html.TextNode
)

type (
	Attribute = html.Attribute
	Node      = html.Node
)

func CopyNode(in *Node) (out *Node) {
	out = &Node{
		Attr:      make([]Attribute, len(in.Attr)),
		DataAtom:  in.DataAtom,
		Data:      in.Data,
		Namespace: in.Namespace,
		Type:      in.Type,
	}
	for i := 0; i < len(in.Attr); i++ {
		out.Attr[i] = in.Attr[i]
	}
	if IsWhitespace(out) {
		out.Data = consecutiveNewlines.ReplaceAllString(out.Data, "\n")
	}
	DebugNodeOpen(out)
	for n := in.FirstChild; n != nil; n = n.NextSibling {
		out.AppendChild(CopyNode(n))
	}
	DebugNodeClose(out)
	return
}

func NewNode(tag atom.Atom, attr ...string) (n *Node) {
	n = &Node{
		DataAtom: tag,
		Data:     tag.String(),
		Type:     ElementNode,
	}
	if len(attr)%2 != 0 {
		panic(fmt.Sprintf("attribute key given without a value: %v", attr))
	}
	for i := 0; i < len(attr); i += 2 {
		n.Attr = append(n.Attr, Attribute{
			Key: attr[i],
			Val: attr[i+1],
		})
	}
	return
}

func NodeTypeString(t html.NodeType) string {
	switch t {
	case CommentNode:
		return "CommentNode"
	case DoctypeNode:
		return "DoctypeNode"
	case DocumentNode:
		return "DocumentNode"
	case ElementNode:
		return "ElementNode"
	case ErrorNode:
		return "ErrorNode"
	case RawNode:
		return "RawNode"
	case TextNode:
		return "TextNode"
	default:
		return fmt.Sprint(t)
	}
}

var consecutiveNewlines = regexp.MustCompile("\n\n+")
