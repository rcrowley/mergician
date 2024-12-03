package html

import (
	"fmt"
	"html"
	"regexp"
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
	m := regexp.MustCompile(`^(<[^>]+>)\s*(\+?=)\s*(<[^>]+>)$`).FindStringSubmatch(s)
	//log.Printf("%#v", m)
	if len(m) != 4 {
		err = RuleError(fmt.Sprintf("invalid rule: %s", s))
		return
	}
	if r.Dst, err = ParseString(m[1]); err != nil {
		return
	}
	//log.Print(String(r.Dst))
	r.Op = m[2]
	//log.Print(r.Op)
	if r.Src, err = ParseString(m[3]); err != nil {
		return
	}
	//log.Print(String(r.Src))
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

type RuleError string

func (err RuleError) Error() string {
	return string(err)
}

type Rules []Rule

func (rs *Rules) Set(s string) error {
	r, err := ParseRule(s)
	if err != nil {
		return err
	}
	*rs = append(*rs, r)
	return nil
}

func (rs *Rules) String() string {
	ss := make([]string, len(*rs))
	for i, r := range *rs {
		ss[i] = r.String()
	}
	return strings.Join(ss, ", ")
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
