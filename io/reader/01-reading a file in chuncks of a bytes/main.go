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
