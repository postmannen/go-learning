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
	return http.HandlerFunc(fn)
}

func main() {
	http.HandleFunc("/", testHandler("apekatt"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error: Starting web listener: ", err)
	}
}
