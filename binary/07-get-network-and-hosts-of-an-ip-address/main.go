package main

import (
	"fmt"
)

func main() {
	//The 32 bit unsigned int representing an ip address.
	var addr uint32 = 0xC0A87864 //192.168.120.100
	fmt.Printf("addr = %b\n", addr)

	//The 32 bit unsigned int representing an ip mask.
	var mask uint32 = 0xFFFFFF80
	var inverseMask uint32 = ^mask
	fmt.Printf("mask=%032b, inverse mask=%032b\n", mask, inverseMask)

	//Split up the uint32 into 4 bytes, where each byte represent an element in a slice.
	adr := readOctets(addr)
	msk := readOctets(mask)

	fmt.Println(getPrefix(adr, msk))
	fmt.Println(getHosts(adr, msk))
}

//getPrefix will get the network portion of the address.
// The ^ operator flips all the bits to it's opposite value,
// meaning 0 becomes 1, and 1 becomes a 0.
func getPrefix(a []byte, m []byte) []byte {
	v := make([]byte, 4)
	for i := 0; i < 4; i++ {
		//Do an AND operation with the mask on the address.
		v[i] = a[i] & m[i]
	}
	return v
}

//getHosts will get the hosts portion of the address.
func getHosts(a []byte, m []byte) []byte {
	v := make([]byte, 4)
	for i := 0; i < 4; i++ {
		//Do an and operation with the inverse mask on the address.
		v[i] = a[i] & ^m[i]
	}
	return v
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
