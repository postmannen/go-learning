/*
Same as test nr 19 for testing chaining functions on a struct,
but this one returns an interface from the methods instead of
a pointer to the structs, so we only make the methods available
and not the types of the structs while chaining.
*/
package main

import (
	"fmt"
	"strings"
)

type aString struct {
	text string
}

type stringer interface {
	lower() stringer
	upper() stringer
}

func (a *aString) lower() stringer {
	a.text = strings.ToLower(a.text)
	return a
}

func (a *aString) upper() stringer {
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
