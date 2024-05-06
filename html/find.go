package html

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

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
