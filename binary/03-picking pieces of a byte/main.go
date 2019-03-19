package main

import (
	"fmt"
)

func main() {
	var a uint16
	a = 0xF8 //11111000
	fmt.Printf("The original byte : %b\n", a)

	fmt.Printf("The first 4 bits (nibble) : %b\n", a>>4)
	a &= 0xF
	fmt.Printf("The last 4 bits (nibble) : %b\n", a)
}
