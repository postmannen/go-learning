package main

import (
	"bytes"
	"fmt"
)

func main() {

	byteSlice1 := []byte{1, 2, 3, 4, 5}
	byteSlice2 := []byte{11, 12, 13, 14, 15}

	//this one createa a buffer of type bytes.Buffer
	var myBuf1 bytes.Buffer
	myBuf1.Write(byteSlice1)
	fmt.Printf("myBuf1 contains = %v, and the type=%T\n", myBuf1, myBuf1)

	//this one creates a buffer of type *bytes.Buffer
	myBuf2 := bytes.NewBuffer(byteSlice1)
	fmt.Printf("myBuf2 contains = %v, and the type=%T\n", myBuf2, myBuf2)
	myBuf2.Write(byteSlice2)

	//ReadByte will read one byte from the start of the buffer,
	//and remove the read byte from the buffer
	myByte, err := myBuf2.ReadByte()
	if err != nil {
		fmt.Println("Error: Reading byte from buffer: ", err)
	}
	fmt.Println("myByte = ", myByte)

	//ReadByte will read the next byte from the start of the buffer,
	//and remove the read byte from the buffer
	myByte, err = myBuf2.ReadByte()
	if err != nil {
		fmt.Println("Error: Reading byte from buffer: ", err)
	}
	fmt.Println("myByte = ", myByte)

	//range the remaining items of the buffer, and print them to the screen
	//will give error = EOF when done reading the buffer
	for {
		b, err := myBuf2.ReadByte()
		if err != nil {
			fmt.Println("For loop Error :", err)
			break
		}
		fmt.Println("For loop read from buffer: ", b)
	}
}
