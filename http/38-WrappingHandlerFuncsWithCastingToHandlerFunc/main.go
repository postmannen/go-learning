//Test, converting a function with a handler signature into a HandlerFunc type.
package main

import (
	"fmt"
	"net/http"
)

func testHandler(s string) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The string passed in was : ", s)
	}
	//The conversion is not needed since the signature of the function is matching the signature
	// needed for the http.HandlerFunc function type, but it was added in this example for
	// testing and learning.
	return http.HandlerFunc(fn)
}

func main() {
	http.HandleFunc("/", testHandler("apekatt"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error: Starting web listener: ", err)
	}
}
