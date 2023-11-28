package html

import (
	"bytes"
	"io"
	"os"

	"golang.org/x/net/html"
)

// Print serializes and prints the *Node to standard output.
func Print(n *Node) (err error) {
	err = Render(os.Stdout, n)
	return
}

// Render is an alias for x/net/html's Render function.
func Render(w io.Writer, n *Node) error {
	return html.Render(w, n)
}

// String renders the *Node to a string and returns it. In case of error,
// the return value is the error string instead. If handling this error
// is important to you, use Render instead.
func String(n *Node) string {
	var b bytes.Buffer
	err := html.Render(&b, n)
	if err != nil {
		return err.Error()
	}
	return b.String()
}
