package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, `Usage: cms
`)
	}
	flag.Parse()
}
