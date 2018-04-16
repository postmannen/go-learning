package main

import (
	"io"
)

//create a struct type with a field that accepts any kind of type
//that satisfies the io.Reader interface.
type myType struct {
	myField io.Reader
}

func (m myType) Read(p []byte) (n int, err error) {
	//lets use io.Readers Read method, and embed it into myType's Read method.
	n, err = m.myField.Read(p)
	return n, err
}

func readAll() {

}

func main() {

}
