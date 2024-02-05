package html

import (
	"encoding/json"
	"strings"
)

// TextOnlyNode could be structured as a tree of linked lists, like the DOM,
// but if we do that then the somewhat naive indexing libraries won't be able
// to follow. So we have to structure the tree using slices, which reflection,
// JSON encoding, etc. can follow.
type TextOnlyNode struct {
	Nodes []TextOnlyNode
	Text  string
}

func Text(in *Node) (out TextOnlyNode) {
	for i := in.FirstChild; i != nil; i = i.NextSibling {
		o := Text(i)
		if len(o.Nodes) > 0 || o.Text != "" {
			out.Nodes = append(out.Nodes, o)
		}
	}
	if in.Type == TextNode && strings.TrimSpace(in.Data) != "" {
		out.Text = in.Data
	}
	return
}

func (n TextOnlyNode) String() string {
	b, err := json.MarshalIndent(n, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(b)
}
