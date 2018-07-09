/*
	Test out embedding an existing type, and inherit all the methods that the
	original type have. This is done by closing the underlying type into struct{T}
*/
package main

import (
	"bytes"
	"fmt"
	"log"
)

var oneBuffer bytes.Buffer

//here we don't inherit the methods of bytes.Buffer since we dont specify it like
//type anotherBuffer struct{bytes.Buffer}
type anotherBuffer bytes.Buffer

var twoBuffer anotherBuffer

//by using struct{bytes.Buffer} when initializing we can access all of the the underlying
//methods of bytes.Buffer in our new type yetAnother buffer.
type yetAnotherBuffer struct{ bytes.Buffer }

//ReadPrinter is a new method added to yetAnotherPrinter to test out creating new
//methods to new types which are embedding another allready existing type.
//So here we should get all of bytes.Buffer's methods which yetAnotherBuffer is
//embeddig, and we should also get the extra method readPrinter
func (y yetAnotherBuffer) readPrinter() (n int, err error) {
	for {
		readVar := make([]byte, 4)
		n, err = y.Read(readVar)
		if n == 0 {
			log.Println("info : reached the end of reading the buffer")
			return n, err
		}
		if err != nil && n != 0 {
			log.Println("error : could not read the whole buffer")
			return n, err
		}
		fmt.Println(string(readVar))
	}

}

var threeBuffer yetAnotherBuffer

func main() {
	fmt.Println("--------------------------------TEST1-----------------------------------------")
	//threeBuffer is of type  yetanotherBuffer is just a new type embedding bytes.Buffer
	//since we used the syntax struct{bytes.Buffer} when declaring the type we also inherit
	//the methods of bytes.Buffer, and we can for example use it's Write method
	n, err := threeBuffer.Write([]byte("this is yet another test string"))
	if err != nil {
		log.Println("error : buffer.write : ", err)
	}

	n, err = threeBuffer.readPrinter()
	if err != nil {
		log.Println("error : threeBuffer.readPrinter : ", err)
	}

	//----------------- Normal write and read operations ---------------------------
	fmt.Println("--------------------------------TEST2-----------------------------------------")
	n, err = oneBuffer.Write([]byte("This is a test string"))
	if err != nil {
		log.Println("error : buffer.write : ", err)
	}
	fmt.Println("characters written = ", n)

	//Read from the buffer into 'readVar'
	for {
		readVar := make([]byte, 4)
		n, err = oneBuffer.Read(readVar)

		if n == 0 {
			log.Println("info : reached the end of reading the buffer")
			break
		}
		if err != nil && n != 0 {
			log.Println("error : could not read the whole buffer")
			break
		}
		fmt.Println(string(readVar))
	}

}
