package html

import (
	"fmt"
	"html"
	"strings"

	"golang.org/x/net/html/atom"
)

func DefaultRules() []Rule {
	return []Rule{
		// TODO <title>
		// TODO <head>
		{NewNode(atom.Article, "class", "body"), "=", NewNode(atom.Body)},
		{NewNode(atom.Div, "class", "body"), "=", NewNode(atom.Body)},
		{NewNode(atom.Section, "class", "body"), "=", NewNode(atom.Body)},
	}
}

type Rule struct {
	Dst *Node
	Op  string // panics if not "=" or "+="
	Src *Node
}

func (r Rule) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		nodeStringForRule(r.Dst),
		r.Op,
		nodeStringForRule(r.Src),
	)
}

func nodeStringForRule(n *Node) string {
	ss := make([]string, len(n.Attr)+3)
	ss[0] = "<"
	ss[1] = n.DataAtom.String()
	for i, attr := range n.Attr {
		ss[i+2] = fmt.Sprintf(" %s=%q", html.EscapeString(attr.Key), html.EscapeString(attr.Val))
	}
	ss[len(n.Attr)+2] = ">"
	return strings.Join(ss, "")
}
