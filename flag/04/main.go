package main

import (
	"flag"
	"fmt"
	"strings"
)

type hobbies []string

func (h *hobbies) String() string {
	return fmt.Sprint(*h)
}

func (h *hobbies) Set(value string) error {
	for _, hobby := range strings.Split(value, ",") {
		*h = append(*h, hobby)
	}
	return nil
}

func main() {
	var hobbiesFlag hobbies
	flag.Var(&hobbiesFlag, "hobbies", "comma separated list of hobbies")

	// Enable command-line parsing
	flag.Parse()

	fmt.Printf("Your hobbies are: ")
	for _, hobby := range hobbiesFlag {
		fmt.Printf("%s ", hobby)
	}
	fmt.Println()
}