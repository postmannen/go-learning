/*
	Example for picking out only the value of the 4 least significant bits of a byte,
	use that value to lookup the map.
*/
package main

import "fmt"

type binaryValue uint8

const (
	null  binaryValue = iota //00
	one                      //01
	two                      //10
	three                    //11
)

var m map[binaryValue]string = map[binaryValue]string{
	0: "null",
	1: "one",
	2: "two",
	3: "three",
}

func main() {

	//we have a byte 82
	var b uint8 = 82 //10000010 which is the numeric value 2

	//we want to check only the 4 least significant bits
	v := b & 0xA
	fmt.Printf("The 4 lsb are binary=%b, numeric=%v\n", v, v)

	fmt.Printf("Using this number as key in the map returns the value = %v\n", m[binaryValue(v)])
}
