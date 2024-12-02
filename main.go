package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rcrowley/mergician/files"
	"github.com/rcrowley/mergician/html"
)

func init() {
	log.SetFlags(0)
}

func main() {
	output := flag.String("o", "-", "write to this file instead of standard output")
	rules := new(html.Rules)
	flag.Var(rules, "r", "use a custom rule for merging inputs (overrides all defaults; may be repeated)")
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: mergician [-o <output>] [-r <rule>[...]] <input>[...]
  -o <output>   write to this file instead of standard output
  -r <rule>     use a custom rule for merging inputs (overrides all defaults;
                may be repeated)
                each rule is a destination HTML tag with optional attributes,
                "=" or "+=", and a source HTML tag with optional attributes
                default rules: <article class="body"> = <body>
                               <div class="body"> = <body>
                               <section class="body"> = <body>
  <input>[...]  one or more input HTML, Markdown, or Google Doc HTML-in-zip files
`)
	}
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("need at least one input HTML, Markdown, or Google Doc HTML-in-zip file")
	}
	in := must2(files.ParseSlice(flag.Args()))

	// Short-circuit if we've only been given one argument. We check this after
	// the call to files.ParseSlice() for the side-effect of files.ParseSlice()
	// rendering Markdown to HTML. If there's only one file, though, let's not
	// try to merge it with nothing, which will superficially change the HTML
	// and change the hashes stored alongside the HTML rendered from Markdown.
	if flag.NArg() == 1 {
		if *output == "-" {
			f := must2(os.Open(fmt.Sprintf("%s.html", strings.TrimSuffix(flag.Arg(0), filepath.Ext(flag.Arg(0))))))
			defer f.Close()
			must2(io.Copy(os.Stdout, f))
		}
		return
	}

	if len(*rules) == 0 {
		*rules = html.DefaultRules()
	}

	out := must2(html.Merge(in, *rules))

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
