// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:] //ta med argumentene etter programnavnet, og legg i 'files'
	if len(files) == 0 { //hvis det ikke er noen argumenter
		countLines(os.Stdin, counts) //kjør countlines med peker til os.Stdin som input
	} else {
		for _, arg := range files { //loop over alle argumentene, og legg filnavn i 'arg'
			f, err := os.Open(arg) //åpne fil og legg peker til fil i 'f'
			if err != nil {        //hvis feilmelding er annet enn 0
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue //kjør neste iteration
			}
			countLines(f, counts) //kjør countlines med peker f til fil som input
			f.Close()             //lukk fil
		}
	}
	for line, n := range counts { //range over map[string]int counts
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
