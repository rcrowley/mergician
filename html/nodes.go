package html

import (
	"fmt"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Node = html.Node

func NewNode(tag atom.Atom, attr ...string) (n *Node) {
	n = &Node{
		DataAtom: tag,
		Data:     tag.String(),
		Type:     html.ElementNode,
	}
	if len(attr)%2 != 0 {
		panic(fmt.Sprintf("attribute key given without a value: %v", attr))
	}
	for i := 0; i < len(attr); i += 2 {
		n.Attr = append(n.Attr, html.Attribute{
			Key: attr[i],
			Val: attr[i+1],
		})
	}
	return
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
