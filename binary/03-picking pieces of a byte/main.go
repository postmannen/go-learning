/*
	This example shows how to bit shift out specified number
	of bits out of some value.
*/

package main

import (
	"fmt"
)

func main() {
	var a uint16
	a = 0xF8 // in binary 11111000
	fmt.Printf("The original byte : %b\n", a)

	fmt.Printf("The first 4 bits (nibble) : %b\n", a>>4)
	a &= 0xF
	fmt.Printf("The last 4 bits (nibble) : %b\n", a)
}
