package main

import (
	"flag"
	"fmt"
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
	in, err := parse(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
	out, err := html.Merge(in, html.DefaultRules())
	if err != nil {
		log.Fatal(err)
	}

	if *output == "-" {
		err = html.Print(out)
	} else {
		err = html.RenderFile(*output, out)
	}
	if err != nil {
		log.Fatal(err)
	}
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
