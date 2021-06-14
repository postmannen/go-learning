/*
io.Reader is an interface of type Reader, that is defined within the io package.
The Reader interface only has one description of a method defined, and that is:
	type Reader interface {
		read(p []byte) (n int, err error)
	}
This means that any other type that has a "read(p []byte) (n int, err error)" method satisfies the Reader interface.
When a type satisfy an interface it will also become the type interface it satisfies. That means if we have "type badger string",
 with a read method, the 'badger' type will also become a type 'Reader'.
This lets us use the 'badger' type as input or output anywhere an io.Reader is wanted.
We can also create a new variable of type 'Reader' and fill that with a type 'badger' value.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	//open a file, and put the handle in fp. fp becomes a pointer to the file
	fp, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("ERROR: opening file : %v\n", err)
	}
	fmt.Printf("fp = %v, and the type = %T\n", fp, fp)

	//create a slice of byte with size 1 byte.
	//The reason for creating a slice is that all Readers and Writers take a slice of byte (b []byte)
	data := make([]byte, 1)

	//create an endless for loop. Will 'break' of err !=nil, which means f.ex. eof
	for {
		numRead, err := fp.Read(data)
		if err != nil {
			fmt.Printf("ERROR: reading file : %v\n", err)
			//will break the loop if an actual error, or a eof is received
			break
		}

		fmt.Printf("data = %v : type = %T : numRead = %v\n", string(data), data, numRead)
	}
}
