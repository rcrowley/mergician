package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rcrowley/mergician/html"
	"github.com/rcrowley/mergician/markdown"
)

func init() {
	log.SetFlags(0)
}

func main() {
	output := flag.String("o", "-", "write to this file instead of standard output")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: mergician [-o <output>] <input>[...]
  -o <output>   write to this file instead of standard output
  <input>[...]  pathname to one or more input HTML or Markdown files
`)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("need at least one input HTML or Markdown file")
	}
	in := must2(parse(flag.Args()))

	// Short-circuit if we've only been given one argument. We check this after
	// the call to parse() for the side-effect of parse() rendering Markdown to
	// HTML. If there's only one file, though, let's not try to merge it with
	// nothing, which will superficially change the HTML and change the hashes
	// stored alongside the HTML rendered from Markdown.
	if flag.NArg() == 1 {
		if *output == "-" {
			f := must2(os.Open(fmt.Sprintf("%s.html", strings.TrimSuffix(flag.Arg(0), filepath.Ext(flag.Arg(0))))))
			defer f.Close()
			must2(io.Copy(os.Stdout, f))
		}
		return
	}

	out := must2(html.Merge(in, html.DefaultRules()))

	if *output == "-" {
		must(html.Print(out))
	} else {
		must(html.RenderFile(*output, out))
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func must2[T any](v T, err error) T {
	must(err)
	return v
}

func parse(pathnames []string) (in []*html.Node, err error) {
	in = make([]*html.Node, len(pathnames))
	for i, pathname := range pathnames {

		if ext := filepath.Ext(pathname); ext == ".md" {
			d, err := markdown.ParseFile(pathname)
			if err != nil {
				return nil, err
			}
			pathname = fmt.Sprintf("%s.html", strings.TrimSuffix(pathname, ext))
			if err := markdown.RenderFile(pathname, d); err != nil {
				return nil, err
			}
		}

		if in[i], err = html.ParseFile(pathname); err != nil {
			return nil, err
		}

	}
	return in, nil
}
