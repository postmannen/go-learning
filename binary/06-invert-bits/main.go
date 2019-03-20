/*
	Invert the bits in an uint32.
	Make all 0's to 1's, and all 1's to 0's.
	It is the ^ operator who do the inversion.
*/
package main

import (
	"fmt"
)

func main() {
	var addr uint32 = 0xC0A87864 //192.168.120.100
	fmt.Printf("byte b = %b\n", addr)

	var mask uint32 = 0xFFFFFF00
	var inverseMask uint32 = ^mask
	fmt.Printf("mask m=%032b, mi=%032b\n", mask, inverseMask)

	netAddr := addr & mask
	fmt.Printf("netAddr=%032b\n", netAddr)
}
