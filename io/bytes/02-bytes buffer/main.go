package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {

	byteSlice1 := []byte{1, 2, 3, 4, 5}
	byteSlice2 := []byte{11, 12, 13, 14, 15}
	byteSlice3 := []byte{110, 120, 130, 140, 150}

	//this one createa a buffer of type bytes.Buffer
	var myBuf1 bytes.Buffer
	myBuf1.Write(byteSlice1)
	fmt.Printf("myBuf1 contains = %v, and the type=%T\n", myBuf1, myBuf1)

	fmt.Println("--------------------------------------")

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

	fmt.Println("--------------------------------------")

	//make a new buffer and write 3 different slices of bytes to it
	myBuf3 := bytes.NewBuffer(byteSlice1)
	myBuf3.Write(byteSlice2)
	fmt.Printf("myBuf3 contains = %v, and the type=%T\n", myBuf3, myBuf3)
	_, err = myBuf3.Write(byteSlice3)
	if err != nil {
		fmt.Println("Error: Write: ", err)
	}
	fmt.Printf("myBuf3 contains = %v, and the type=%T\n", *myBuf3, myBuf3)

	//read the whole buffer to outSlice as a []byte
	outSlice, err := ioutil.ReadAll(myBuf3)
	if err != nil {
		fmt.Println("Error: ioutil.ReadAll: ", err)
	}
	fmt.Println("Outslice contains = ", outSlice)
}
