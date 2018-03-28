/*
ioutil testing
ioutil is a package with some functions to make file reading and writing easier.
The functions are just using os.Open and Read functions.
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	rh, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("error opening file: ", err)
		os.Exit(1)
	}
	defer rh.Close()

	/*
		ReadAll will take io.Reader as input. And since rh which is a *os.File have a Read method,
		it also fullfills the requirement of the Reader interface, and is by that an io.Reader
		ReadAll will read the whole input, and put it in a []byte
	*/
	fileData, err := ioutil.ReadAll(rh)
	if err != nil {
		fmt.Println("error reading file: ", err)
	}
	fmt.Println(string(fileData))
}
