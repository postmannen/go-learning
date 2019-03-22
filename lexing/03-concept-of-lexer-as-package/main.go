// The idea here is to learn and test out the concept of having local variables and methods in a package (lexml),
// and only expose what we want to the caller (main.go).
//
package main

import (
	"fmt"

	"github.com/postmannen/go-learning/lexing/03-concept-of-lexer-as-package/lexml"
)

var aFile = []string{"one line", "two line", "three line", "four line", "five line", "six line", "seven line"}

func main() {
	ch := lexml.StartLexing(aFile)

	for v := range ch {
		fmt.Println("*** Got from channel : ", v)
	}
}
