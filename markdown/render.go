package markdown

import (
	"io"
	"os"
)

// Print serializes and prints the *Markdown to standard output.
func Print(d *Document) error {
	return Render(os.Stdout, d)
}

// Render renders the *Markdown to HTML and writes it to the io.Writer.
func Render(w io.Writer, d *Document) error {
	_, err := w.Write(d.Bytes())
	return err
}

// String renders the *Markdown to a string and returns it. In case of error,
// the return value is the error string instead. If handling this error
// is important to you, use Render instead.
func String(d *Document) string {
	return d.String()
}
