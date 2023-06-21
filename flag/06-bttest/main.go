package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var version = "0.1"

const usage = ` Usage:
	-version				prints version

`

func main() {
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%v", usage) }

	var versionFlag bool

	flag.BoolVar(&versionFlag, "version", false, "print version")
	flag.Parse()

	log.Printf(" * info: flag.NArg: %v\n", flag.NArg())
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "%v", usage)
	}

	if versionFlag {
		fmt.Printf(" * version: %v\n", version)
	}

}
