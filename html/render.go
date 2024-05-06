package html

import (
	"bytes"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

// Print serializes and prints the *Node to standard output.
func Print(n *Node) error {
	return Render(os.Stdout, n)
}

// Render is almost an alias for x/net/html's Render function but ensures
// files end with a trailing '\n' character.
func Render(w io.Writer, n *Node) error {
	if err := html.Render(w, n); err != nil { // TODO make it stop with the XML-style self-closing tags; I hate that
		return err
	}
	_, err := w.Write([]byte{'\n'})
	return err
}

// RenderFile writes an HTML document to a file, overwriting if the file
// already exists.
func RenderFile(pathname string, n *Node) error {
	f, err := os.Create(pathname)
	if err != nil {
		return err
	}
	defer f.Close()
	return Render(f, n)
}

// String renders the *Node to a string and returns it. In case of error,
// the return value is the error string instead. If handling this error
// is important to you, use Render instead.
func String(n *Node) string {
	var b bytes.Buffer
	err := Render(&b, n)
	if err != nil {
		return err.Error()
	}
	return b.String()
}

func printNodeAsTree(n *Node, indent string) {
	return
	switch n.Type {
	case html.ElementNode:
		log.Printf("%s<%s>\n", indent, n.DataAtom) // TODO n.Attr
	case html.TextNode:
		log.Printf("%s%q\n", indent, n.Data)
	default:
		log.Printf("%s%s %+v\n", indent, NodeTypeString(n.Type), n)
	}
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		printNodeAsTree(child, indent+"\t")
	}
	switch n.Type {
	case html.ElementNode:
		log.Printf("%s</%s>\n", indent, n.DataAtom)
	}
}
