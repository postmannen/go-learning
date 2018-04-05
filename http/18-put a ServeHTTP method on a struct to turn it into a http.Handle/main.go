/*
Want to test creating a method to a struct that satisfies the http.Handler interface
*/
package main

import (
	"fmt"
	"net/http"
)

type someStruct struct {
	someString string
}

func (s *someStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.someString))
}

func main() {
	http.Handle("/", &someStruct{someString: "This is a little test of filling a string inside a struct from a ServeHTTP method on a struct :-)"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("ListenAndServe failed with error : ", err)
	}

}
