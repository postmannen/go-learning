/*
	Calculates the Prefix, and the address range for that prefix.
	Using bitshift and logical AND (&) to split up the uint32 to uint8.

	Using logical AND (&) to get the network prefix,
	and logical XOR (^) to get the inverse value of the mask to calulate
	the range of addresses for that prefix.
*/
package main

import (
	"fmt"
)

func main() {
	//The 32 bit unsigned int representing an ip address.
	var addr uint32 = 0xC0A87880
	fmt.Printf("addr = %b, human readable = %v\n", addr, convertToOctets(addr))

	//The 32 bit unsigned int representing an net mask.
	var mask uint32 = 0xFFFFFFE0
	fmt.Printf("mask=%032b, human readable = %v\n", mask, convertToOctets(mask))

	fmt.Println("-----------Get Prefix----------------------------------------------------------")
	prefix := getPrefix(addr, mask)
	fmt.Printf("The prefix=%032b, human readable=%v\n", prefix, convertToOctets(prefix))

	fmt.Println("-----------Get Hosts----------------------------------------------------------")
	fmt.Println("Inverse mask = ", ^mask)
	maxHosts := prefix + ^mask
	fmt.Printf("starting at = %v, ending at = %v\n", convertToOctets(prefix), convertToOctets(maxHosts))
}

//getPrefix will get the network portion of the address.
// The ^ operator flips all the bits to it's opposite value,
// meaning 0 becomes 1, and 1 becomes a 0.
func getPrefix(addr uint32, mask uint32) uint32 {
	prefix := addr & mask
	//fmt.Printf("%032b, %032b, %032b\n", addr, mask, prefix)
	return prefix
}

//readOctets will read chuncs of 8 bits from the uint32 address,
// and put each byte into a slice which will be returned from the function.
func convertToOctets(addr uint32) []byte {
	var numBits uint32 = 0xFF
	bSlice := make([]byte, 4)

	for i := 3; i >= 0; i-- {
		//Take out the 8 least significant bits (lsb) by doing an and operation
		// with the binary value 11111111 which eqals FF in hexadecimal,
		// and append those 8 to the slice.
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
