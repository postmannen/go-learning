package main

import (
	"fmt"
)

type item int

const (
	startTag item = iota
	endTag
	argumentKey
	argumentValue
)

func main() {
	foundWhileLexing := argumentKey

	if foundWhileLexing == argumentKey {
		fmt.Printf("Value = %v, and the type = %T\n", foundWhileLexing, foundWhileLexing)
	}
}
