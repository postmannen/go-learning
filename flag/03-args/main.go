package main

import (
	"flag"
	"fmt"
)

// go run main.go -lines=10 gris apekatt
// If a normal flag is given as the first argument, then what
// comes after will be Args.
//
// go run main.go apekatt -lines=10 gris
// If an Arg is given as the first argument all that comes after
// is also taken as Args, even the normal flag with a dash infront
// becomes an Arg.

func main() {
	var lines int
	flag.IntVar(&lines, "lines", 1, "number of lines")
	flag.Parse()

	if len(flag.Args()) > 0 {
		for _, f := range flag.Args() {
			fmt.Printf("also got arg: %v\n", f)
		}
	}
}
