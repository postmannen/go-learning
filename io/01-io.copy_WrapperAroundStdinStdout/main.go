/*
Example for using io.Copy to copy directly from STDIN which is an io.Reader
to STDOUT which is an io.Writer.
This could be achieved without building the structs and methods used here,
but the idea were to wrap the Reader and the Writer into new methods.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type mySTDIN struct{}

// Read wraps a STDIN reader
func (m *mySTDIN) Read(b []byte) (int, error) {
	fmt.Printf("in >\n")
	i := os.Stdin

	n, err := i.Read(b)
	if err != nil {
		if err == io.EOF {
			log.Println("ctrl+d was pressed, EOF")
		}
		log.Println("Read stdin failed: ", err, " : n=", n)
	}
	fmt.Println("Bytes read = ", n)
	return n, err
}

type mySTDOUT struct{}

// Write wraps a STDOUT writer
func (m *mySTDOUT) Write(b []byte) (int, error) {
	fmt.Printf("out >\n")
	o := os.Stdout

	n, err := o.Write(b)
	if err != nil {
		log.Println("Write stdout failed: ", err, " : n=", n)
	}
	fmt.Println("Bytes written = ", n)
	return n, err
}

func main() {
	in := mySTDIN{}
	out := mySTDOUT{}

	// io.Copy will continue to copy bytes from the Reader to the Writer until io.EOF
	// is received. In the terminal shell ctrl+d will give the io.EOF.
	if _, err := io.Copy(&out, &in); err != nil {
		log.Println("io.Copy failed: ", err)
	}

}
