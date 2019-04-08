package main

import "fmt"

func main() {
	//tmp = &bytes.Buffer{}
	//binary.Write(tmp, binary.LittleEndian, int8(b.Pcmd.Yaw))
	//cmd.Write(tmp.Bytes())

	number := int8(-4)
	fmt.Printf("Before casting into the slice of bytes: %#v, %T\n", number, number)

	bSlice := []byte{byte(number)}
	fmt.Printf("Casting into byte, and putting on a slice of bytes: %#v\n", bSlice)

	//then cast it back....
	fmt.Printf("Casting it back to int8 directly in print: %#v, %T\n", int8(bSlice[0]), int8(bSlice[0]))
	//and try putting it in another variable
	number2 := int8(bSlice[0])
	fmt.Printf("Casting it back to int8, and assigning to variable: %#v, %T\n", number2, number2)

	/*
		The print statements will give the following output:
			Before casting into the slice of bytes: 				-4, int8
			Casting into byte, and putting on a slice of bytes: 	[]byte{0xfc}
			Casting it back to int8 directly in print: 				-4, int8
			Casting it back to int8, and assigning to variable: 	-4, int8

		int8 can store values from -127 -> 127
		If we convert this to a byte the lover 128 values are the positive numbers,
		and the upper 128 values are the negative numbers
	*/
}
