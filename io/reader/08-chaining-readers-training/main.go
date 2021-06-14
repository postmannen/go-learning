package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type name struct {
	names io.Reader
}

func (na name) Read(p []byte) (n int, err error) {
	fmt.Println("using the names.Read method")
	n, err = na.names.Read(p)
	return n, err
}

//Take an input of type io.Reader, and return a *name with the content of the io.Reader
func newName(p io.Reader) *name {
	return &name{names: p}
}

func namePrinter(p io.Reader) {
	names := []byte{}
	//loop over the input, read a byte chunc and append to names variable
	for {
		//make a bugger of []byte with the length of 4
		buf := make([]byte, 4)
		_, err := p.Read(buf)
		if err != nil {
			fmt.Println("finnished reading with namePrinter function")
			break
		}
		names = append(names, buf...)
	}
	fmt.Printf("namePrinter: content read = %v \n", names)
	fmt.Printf("namePrinter: type of data read = %T\n", p)
}

func main() {
	//test1
	students := name{names: strings.NewReader("ole,eva,knut")}
	buf := make([]byte, 256)
	students.Read(buf)
	fmt.Printf("buffer contains = %v\n", string(buf))

	//test2
	fh, err := os.Open("names.txt")
	if err != nil {
		fmt.Println("error: opening file = ", err)
	}
	defer fh.Close()

	students2 := newName(fh)
	buf2 := make([]byte, 256)
	students2.Read(buf2)
	fmt.Printf("buf2 contains = %v\n", string(buf2))

	fmt.Println("---------------------------------TEST3----------------------------")
	namePrinter(students2)
}
