/*
Want to test creating a method to a struct that satisfies the http.Handler interface
*/
package main

import (
	"fmt"
	"net/http"
	"sync"
)

type someStruct struct {
	someString string
	once       sync.Once
}

func (s *someStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(s.someString))

	/*
		putting a 'sync.Once' type field in the struct gives us the possibility to only call a function once.
		Here the function with the write method below will only be called the first time a web page is opened. The second opening of the
		same page will not show the text.
	*/
	s.once.Do(func() {
		w.Write([]byte("This should be written only the first time"))
	})
}

func main() {
	http.Handle("/", &someStruct{someString: "This is a little test of filling a string inside a struct from a ServeHTTP method on a struct :-)\n"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("ListenAndServe failed with error : ", err)
	}

}
