package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
  <input>[...]  pathname to one or more input HTML files
`)
	}
	flag.Parse()

	var (
		err error
		out *html.Node
	)

	if flag.NArg() == 0 {
		log.Fatal("need at least one input HTML file")
	}
	in := make([]*html.Node, flag.NArg())
	for i, pathname := range flag.Args() {
		if in[i], err = html.ParseFile(pathname); err != nil {
			log.Fatal(err)
		}
	}
	out, err = html.Merge(in, html.DefaultRules())

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
