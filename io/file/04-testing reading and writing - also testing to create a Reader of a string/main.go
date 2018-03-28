package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	fmt.Println("-------------------TEST1---------------------")
	//create a file
	fh, err := os.Create("testFile.txt")
	if err != nil {
		fmt.Println("error creating file : ", err)
	}

	fmt.Println("Before the fh.Write, &fh = ", &fh)

	//put some data into it
	fileData := []byte("some data to put in the file\n")
	n, err := fh.Write(fileData)
	if err != nil {
		fmt.Println("error: Writing to file : ", err)
	}
	fmt.Println("Characters written to file = ", n)

	fmt.Println("After the fh.Write, &fh = ", &fh)
	//closing the file
	fh.Close()

	//open the file for reading (using the same filehandle variable)
	fh, err = os.Open("testFile.txt")
	if err != nil {
		fmt.Println("error: opening file for reading : ", err)
	}
	//then read the content of the file. ReadAll takes a io.Reader
	byteSlice1, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println("error: reading file : ", err)
	}

	fmt.Println("Printing out the content read from the file = ", byteSlice1)

	fmt.Println("-------------------TEST2---------------------")

	//create a string, and make it into a type Reader
	myByteSlice := []byte("this is some text")
	myNewReader := strings.NewReader(string(myByteSlice))

	/*
		ReadAll takes an input of type io.Reader.
		Since we've turned the original string 'myByeSlice' into a new variable of type Reader,
		we can now use the string directly in the ReadAll function
	*/
	anotherByteSlice, err := ioutil.ReadAll(myNewReader)
	if err != nil {
		fmt.Println("error: ReadAll : ", err)
	}

	fmt.Println("Printing out the content of the string turned into a Reader = ", string(anotherByteSlice))
}
