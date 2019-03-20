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

	adr := readOctets(addr)
	msk := readOctets(mask)

	getPrefix(adr, msk)
	getHosts(adr, msk)
}

func getPrefix(a []byte, m []byte) {
	for i := 0; i < 4; i++ {
		fmt.Println(a[i] & m[i])
	}

}

func getHosts(a []byte, m []byte) {
	for i := 0; i < 4; i++ {
		fmt.Println(a[i] & ^m[i])
	}
}

func readOctets(addr uint32) []byte {
	var numBits uint32 = 0xFF
	bSlice := make([]byte, 4)

	for i := 3; i >= 0; i-- {
		//Take out the 8 LSB, and append those 8 to the slice.
		b := addr & numBits
		if addr == 0 { //Check if all bits are shifted out.
			break
		}
		bSlice[i] = uint8(b)

		//Shift 8 bits out since we are done with them.
		addr = addr >> 8
	}

	return bSlice
}
