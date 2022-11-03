package main

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) string {
	// var result string
	// for _, v := range s {
	// 	result = string(v) + result
	// }

	b := []byte(s)
	// if len(b) <= 1 {
	// 	return s
	// }

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if !utf8.ValidString(string(b)) {
		return s
	}

	return string(b)
}

func main() {
	s := "Hesten spiser gress og ikke andre dyr som f.eks. apekatter."
	fmt.Printf("%q\n", s)
	fmt.Printf("%q\n", Reverse(s))
	fmt.Printf("%q\n", Reverse(Reverse(s)))

}
