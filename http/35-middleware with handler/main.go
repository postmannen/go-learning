//Using wrapper/middleware with Handler instead of HandlerFunc
package main

import (
	"fmt"
	"net/http"
)

//Test1-------------------------------------------------
type myType struct {
}

//By giving myType a method named ServeHTTP(w,r) it satisfies
//the Handler interface, and then also becomes a type Handler.
func (m myType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The root Handle writes this")
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "The logger wrapper writes this")
		next.ServeHTTP(w, r)
	})
}

func main() {
	a := myType{}

	http.Handle("/", logger(a))

	http.ListenAndServe(":8080", nil)
}
