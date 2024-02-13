package html

import "golang.org/x/net/html/atom"

func FirstParagraph(n *Node) (s string) {
	if p := Find(n, IsAtom(atom.P)); p != nil {
		s = Text(p).String()
	}
	return
}

func Title(n *Node) (s string) {
	if title := Find(n, IsAtom(atom.Title)); title != nil {
		s = Text(title).String()
	}
	return
}
