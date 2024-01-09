package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/rcrowley/mergician/html"
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
		f, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		err = html.Render(f, out)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func parse(pathnames []string) (in []*html.Node, err error) {
	in = make([]*html.Node, len(pathnames))
	for i, pathname := range pathnames {

		if ext := filepath.Ext(pathname); ext == ".md" {
			// TODO render the HTML and hash to disk
			// TODO hashPathname := filepath.Join(filepath.Dir(pathname), fmt.Sprintf(".%s.hash", strings.TrimSuffix(filepath.Base(pathname), ext)))
			// TODO pathname = fmt.Sprintf("%s.html", strings.TrimSuffix(pathname, ext))
		}

		if in[i], err = html.ParseFile(pathname); err != nil {
			return nil, err
		}

	}
	return in, nil
}
