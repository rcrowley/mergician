package html

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func All(funcs ...func(*Node) bool) func(*Node) bool {
	return func(n *Node) bool {
		for _, f := range funcs {
			if !f(n) {
				return false
			}
		}
		return true
	}
}

func Any(funcs ...func(*Node) bool) func(*Node) bool {
	return func(n *Node) bool {
		for _, f := range funcs {
			if f(n) {
				return true
			}
		}
		return false
	}
}

func Find(n *Node, f func(*Node) bool) *Node {
	if f(n) {
		return n
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		if found := Find(child, f); found != nil {
			return found
		}
	}
	return nil
}

func FindAll(n *Node, f func(*Node) bool) (found []*Node) {
	if f(n) {
		found = append(found, n)
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		found = append(found, FindAll(child, f)...) // XXX nil will probably blow up
	}
	return
}

func HasAttr(k, v string) func(*Node) bool {
	return func(n *Node) bool {
		for _, attr := range n.Attr {
			if attr.Namespace == "" && attr.Key == k && attr.Val == v {
				return true
			}
		}
		return false
	}
}

func IsAtom(atoms ...atom.Atom) func(*Node) bool {
	return func(n *Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		for _, a := range atoms {
			if n.DataAtom == a {
				return true
			}
		}
		return false
	}
}

func IsWhitespace(n *Node) bool {
	return n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) == 0
}

func Match(pattern *Node) func(*Node) bool {
	isAtom := IsAtom(pattern.DataAtom)
	hasAttrs := make([]func(*Node) bool, len(pattern.Attr))
	for i, attr := range pattern.Attr {
		hasAttrs[i] = HasAttr(attr.Key, attr.Val)
	}

	return func(n *Node) bool {
		if n.Type != html.ElementNode {
			return false
		}
		if !isAtom(n) {
			return false
		}
		for _, hasAttr := range hasAttrs {
			if !hasAttr(n) {
				return false
			}
		}
		return true
	}
}

func Not(f func(*Node) bool) func(*Node) bool {
	return func(n *Node) bool {
		return !f(n)
	}
}
