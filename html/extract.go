package html

import "golang.org/x/net/html/atom"

func FirstH1(n *Node) string {
	return first(atom.H1, n)
}

func FirstParagraph(n *Node) string {
	return first(atom.P, n)
}

func Title(n *Node) string {
	return first(atom.Title, n)
}

func first(tag atom.Atom, n *Node) (s string) {
	if n = Find(n, IsAtom(tag)); n != nil {
		s = Text(n).String()
	}
	return
}
