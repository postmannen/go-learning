package main

import "fmt"

import "log"

// errorType can be used to cascade several errors together forming
// a chain, and will hold the value of the current error, but can also
// hold the previos error. If no previous error value , use nil.
type errorType struct {
	text     string
	previous error
}

// Error will make the type errorType satisfy the error interface
func (e *errorType) Error() string {
	return e.text
}

// unwrap will first check if the error is of the custom type *errorType,
// and if it is of that type it will call unwrap again on that types value.
func unwrap(err error) {
	switch v := err.(type) {
	case nil:
		log.Println("No error.")
	case *errorType:
		fmt.Println("*** err = ", err)
		if v.previous != nil {
			unwrap(v.previous)
		}
		fmt.Println("No more errors to unwrap, exiting.")
	}

}

func failingFunction1() error {
	err := failingFunction2()
	return &errorType{text: "first error text", previous: err}
}

func failingFunction2() error {
	return &errorType{text: "second error text", previous: nil}
}

func main() {
	err := failingFunction1()

	unwrap(err)
}
