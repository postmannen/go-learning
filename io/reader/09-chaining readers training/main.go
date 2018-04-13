package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type myStructType struct {
	myStructField io.Reader
}

//Creating my own Reader, which under the hood uses io.Reader's Read method.
//So we're chaining two readers, by letting on use the other, and return the
//inner ones output to the outer one.
func (m myStructType) Read(p []byte) (n int, err error) {
	fmt.Println("Using myStructType's Read method")
	return m.myStructField.Read(p)
}

//take an io.Reader as input, and create a new variable as a pointer to a
//memory location, fill that variables field with the io.Reader given as input,
//and return the pointer to the new variable of type myStructType.
//The pointer can then be given to a new variable, and used as a normal variable.
func newMyStructType(p io.Reader) *myStructType {
	return &myStructType{myStructField: p}
}

func printSomething(p io.Reader) {
	myFullRead := []byte{}
	for {
		buf := make([]byte, 4)
		_, err := p.Read(buf)
		if err != nil {
			log.Println("printSomething: Finnished reading")
			break
		}
		myFullRead = append(myFullRead, buf...)
	}
	fmt.Println("printSomething: myFullRead contains = ", string(myFullRead))
}

func main() {
	fmt.Println("-------------------------------------TEST1-----------------------------------------")
	//since we're creating a variable of type myStructType with the new function, we will use
	//myStructType's Read method.
	//The variable will also be of type io.Reader so it can be used where an io.Reader is expected.
	myVariable := newMyStructType(strings.NewReader("Some data to work with since we're creating something here"))
	printSomething(myVariable)

	fmt.Println("-------------------------------------TEST2-----------------------------------------")
	//since we're creating a new type io.Reader variable, and not also making it a myStructType, it
	//will use io.Readers read method, and not myStructType's read method.
	//But since its also of type io.Reader we can use it with the printSomething function.
	myOtherVariable := strings.NewReader("we needed some more data for test2 to work with")
	printSomething(myOtherVariable)
}
