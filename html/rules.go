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

func ParseRule(s string) (r Rule, err error) {
	panic("not implemented")
	return
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

type ruleSlice []Rule

func RuleSlice(rules []Rule) *ruleSlice {
	rs := ruleSlice(rules)
	return &rs
}

func (rs *ruleSlice) Len() int {
	if rs == nil || *rs == nil {
		return 0
	}
	return len(*rs)
}

func (rs *ruleSlice) Set(s string) error {
	r, err := ParseRule(s)
	if err != nil {
		return err
	}
	*rs = append(*rs, r)
	return nil
}

func (rs *ruleSlice) String() string {
	ss := make([]string, len(*rs))
	for i, r := range *rs {
		ss[i] = r.String()
	}
	return strings.Join(ss, ", ")
}
