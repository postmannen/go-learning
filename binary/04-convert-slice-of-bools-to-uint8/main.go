package main

import (
	"fmt"
)

func main() {
	b := []bool{true, false, false, false, true, false, false, true}
	fmt.Println("The values of b which is a []bool, ", b)

	n := convert(b)

	fmt.Printf("Result, binary=%b, numeric=%v\n", n, n)
}

func convert(b []bool) uint8 {
	var n uint8
	for i, v := range b {
		//fmt.Printf("loop, iterating i=%v, value=%v, type=%T\n", i, v, v)

		//If true we swap the least significant bit to 1,
		// if false we leave it to its default which is 0.
		if v == true {
			n = n | 1
		}

		//Bit shift all 1 to the left, which means..we add a zero to the end.
		if i < len(b)-1 {
			n = n << 1
		}
		//fmt.Printf("Value is now, %b\n", n)
	}

	return n
}
