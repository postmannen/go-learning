//Using wrapper/middleware with Handler instead of HandlerFunc
package main

import (
	"fmt"
	"net/http"
)

type myType struct {
}

func (m myType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The root Handle")
}

func logger(next http.Handler) http.Handler {
	//Cast a Handler into a HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the logger wrapper")
		next.ServeHTTP(w, r)
	})
}

func main() {
	a := myType{}

	http.Handle("/", logger(a))
	http.ListenAndServe(":8080", nil)
}
