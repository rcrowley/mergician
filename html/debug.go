package html

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Debug(v ...any) {
	for i := 0; i < len(v); i++ {
		if n, ok := v[i].(*Node); ok && n.Type == ElementNode {
			v[i] = debugElementNode(n)
		}
	}
	if debug {
		log.Print(v...)
	}
}

func Debugf(format string, v ...any) {
	for i := 0; i < len(v); i++ {
		if n, ok := v[i].(*Node); ok && n.Type == ElementNode {
			v[i] = debugElementNode(n)
		}
	}
	if debug {
		log.Printf(format, v...)
	}
}

func DebugNodeOpen(n *Node) {
	if !debug {
		return
	}
	switch n.Type {
	case ElementNode:
		log.Print(debugElementNode(n))
	case TextNode:
		log.Printf("%q\n", n.Data)
	default:
		log.Printf("%s %v\n", NodeTypeString(n.Type), n)
	}
}

func DebugNodeClose(n *Node) {
	if !debug {
		return
	}
	if n.Type == ElementNode {
		log.Printf("</%s>", n.DataAtom)
	}
}

var debug = os.Getenv("MERGICIAN_DEBUG") != ""

func debugElementNode(n *Node) string {
	attrs := make([]string, len(n.Attr))
	for i, attr := range n.Attr {
		attrs[i] = fmt.Sprintf(" %s=%q", attr.Key, attr.Val)
	}
	return fmt.Sprintf("<%s%s>", n.DataAtom, strings.Join(attrs, ""))
}
