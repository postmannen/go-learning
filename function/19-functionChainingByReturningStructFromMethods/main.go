/*
Testing out chaining functions on a struct by returning
a pointer to the structure from each method.
*/
package main

import (
	"fmt"
	"strings"
)

type aString struct {
	text string
}

func (a *aString) lower() *aString {
	a.text = strings.ToLower(a.text)
	return a
}

func (a *aString) upper() *aString {
	a.text = strings.ToUpper(a.text)
	return a
}

func main() {
	a := &aString{text: "MoNkeY"}
	a.lower()
	fmt.Println(a)
	a.upper()
	fmt.Println(a)
	a.lower().upper().lower()
	fmt.Println(a)
}
