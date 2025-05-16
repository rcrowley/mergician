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

func Main(args []string, stdin io.Reader, stdout io.Writer) {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	output := flags.String("o", "-", "write to this file instead of standard output")
	rules := new(html.Rules)
	flags.Var(rules, "r", "use a custom rule for merging inputs (overrides all defaults; may be repeated)")
	flags.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: mergician [-o <output>] [-r <rule>[...]] <input>[...]
  -o <output>   write to this file instead of standard output
  -r <rule>     use a custom rule for merging inputs (overrides all defaults;
                may be repeated);
                each rule is a destination HTML tag with optional attributes,
                "=" or "+=", and a source HTML tag with optional attributes;
                the default rules are: <article class="body"> = <body>
                                       <div class="body"> = <body>
                                       <section class="body"> = <body>
  <input>[...]  one or more input HTML, Markdown, or Google Doc HTML-in-zip documents

Synopsis: mergician merges multiple HTML documents into one based on Microformats-like rules.
`)
	}
	flags.Parse(args[1:])
	if flags.NArg() == 0 {
		log.Fatal("need at least one input HTML, Markdown, or Google Doc HTML-in-zip file")
	}
	in := must2(files.ParseSlice(flags.Args()))

	// Short-circuit if we've only been given one argument. We check this after
	// the call to files.ParseSlice() for the side-effect of files.ParseSlice()
	// rendering Markdown to HTML. If there's only one file, though, let's not
	// try to merge it with nothing, which will superficially change the HTML
	// and change the hashes stored alongside the HTML rendered from Markdown.
	if flags.NArg() == 1 {
		if *output == "-" {
			f := must2(os.Open(fmt.Sprintf("%s.html", strings.TrimSuffix(flags.Arg(0), filepath.Ext(flags.Arg(0))))))
			defer f.Close()
			must2(io.Copy(stdout, f))
		}
		return
	}

	if len(*rules) == 0 {
		*rules = html.DefaultRules()
	}

	out := must2(html.Merge(in, *rules))

	if *output == "-" {
		must(html.Render(stdout, out))
	} else {
		must(html.RenderFile(*output, out))
	}
}

func init() {
	log.SetFlags(0)
}

func main() {
	Main(os.Args, os.Stdin, os.Stdout)
}

func must(err error) {
	if err != nil {
		log.Output(2, err.Error())
	}
}

func must2[T any](v T, err error) T {
	if err != nil {
		log.Output(2, err.Error())
	}
	return v
}
