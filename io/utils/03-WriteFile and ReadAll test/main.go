package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Create a file and put some data into it
	fileData := []byte("this is a text string to put in the file")
	ioutil.WriteFile("testFile.txt", fileData, 0644)

	//open, and read the content of the file
	rh, err := os.Open("testFile.txt")
	if err != nil {
		fmt.Println("error: opening file: ", err)
	}

	//and then read the data
	var myData []byte
	myData, err = ioutil.ReadAll(rh)
	fmt.Println("Read from File = ", string(myData))

}
