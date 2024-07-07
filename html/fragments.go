package html

import "golang.org/x/net/html/atom"

func FirstH1(n *Node) *Node {
	return first(atom.H1, n)
}

func FirstParagraph(n *Node) *Node {
	return first(atom.P, n)
}

func Title(n *Node) *Node {
	return first(atom.Title, n)
}

func first(tag atom.Atom, n *Node) *Node {
	return Find(n, IsAtom(tag))
}
