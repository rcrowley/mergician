package html

import (
	"strings"

	"golang.org/x/net/html/atom"
)

func NewTextNode(s string) *Node {
	return &Node{
		Data: s,
		Type: TextNode,
	}
}

// TextOnlyNode could be structured as a tree of linked lists, like the DOM,
// but if we do that then the somewhat naive indexing libraries won't be able
// to follow. So we have to structure the tree using slices, which reflection,
// JSON encoding, etc. can follow.
type TextOnlyNode struct {
	Nodes []TextOnlyNode
	Text  string
}

func Text(in *Node) (out TextOnlyNode) {
	if in == nil {
		return // makes it safe to do Text(maybeNilNode()).String()
	}
	for i := in.FirstChild; i != nil; i = i.NextSibling {
		o := Text(i)
		if len(o.Nodes) > 0 || o.Text != "" {
			out.Nodes = append(out.Nodes, o)
		}
	}
	switch in.Type {
	case ElementNode:
		switch in.DataAtom {
		case atom.Blockquote, atom.H1, atom.H2, atom.H3, atom.H4, atom.H5, atom.H6, atom.P, atom.Pre:
			out.Text = "\n\n"
		}
	case TextNode:
		if strings.TrimSpace(in.Data) != "" {
			out.Text = in.Data
		}
	}
	return
}

func (n TextOnlyNode) String() string {
	var ss []string
	if n.Text != "" {
		ss = append(ss, n.Text)
	}
	for _, n := range n.Nodes {
		if s := n.String(); s != "" {
			ss = append(ss, s)
		}
	}
	return strings.TrimSpace(strings.Join(ss, " "))
}
