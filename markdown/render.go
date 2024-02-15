package markdown

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/rcrowley/mergician/html"
)

// Print serializes and prints the *Markdown to standard output.
func Print(d *Document) error {
	return Render(os.Stdout, d)
}

// Render renders the *Markdown to HTML and writes it to the io.Writer.
func Render(w io.Writer, d *Document) error {
	b := d.Bytes()
	if _, err := w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<title>`)); err != nil {
		return err
	}

	n, err := html.Parse(d)
	if err != nil {
		return err
	}
	if _, err := w.Write([]byte(html.FirstH1(n))); err != nil {
		return err
	}

	if _, err := w.Write([]byte(`</title>
</head>
<body>
<article>
`)); err != nil {
		return err
	}

	if _, err := w.Write(b); err != nil {
		return err
	}

	if _, err := w.Write([]byte(`</article>
</body>
</html>
`)); err != nil {
		return err
	}
	return nil
}

// RenderFile renders a Markdown document to HTML and writes it to a file,
// only overwriting an existing HTML file if the HTML file's content still
// matches the hash written the last time this Markdown document was rendered
// to this HTML file. If ever we need to upgrade from SHA256, we'll "version"
// the crypto by changing the file extension.
func RenderFile(pathname string, d *Document) error {
	hashPathname := filepath.Join(filepath.Dir(pathname), fmt.Sprintf(".%s.sha256", filepath.Base(pathname)))

	if exists(pathname) {
		if exists(hashPathname) {
			hashHash, err := os.ReadFile(hashPathname)
			if err != nil {
				return err
			}
			htmlHash, err := hashFile(pathname)
			if err != nil {
				return err
			}
			if !bytes.Equal(hashHash, htmlHash) {
				return fmt.Errorf(
					"error: contents of %s does not match %s; not rendering to preserve changes made directly to the HTML",
					pathname,
					hashPathname,
				)
			}
		} else {
			log.Printf(
				"warning: %s exists without %s; changes made directly to the HTML may be lost",
				pathname,
				hashPathname,
			)
		}
	}

	f, err := os.Create(pathname)
	if err != nil {
		return err
	}
	if err := Render(f, d); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	hash, err := hashFile(pathname)
	if err != nil {
		return err
	}
	if err := os.WriteFile(hashPathname, hash, 0666); err != nil {
		return err
	}

	return nil
}

// String renders the *Markdown to a string and returns it. In case of error,
// the return value is the error string instead. If handling this error
// is important to you, use Render instead.
func String(d *Document) string {
	return d.String()
}

func exists(pathname string) bool {
	_, err := os.Stat(pathname)
	return err == nil
}

func hashFile(pathname string) ([]byte, error) {
	f, err := os.Open(pathname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}
	return h.Sum(nil)[:], nil
}
