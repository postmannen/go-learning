package main

import (
	"fmt"
	"net/http"
)

type handlerType string

// Any type that have an ServeHTTP method with the signature below
// will be a valid Handler to be used to serve web pages.
func (h handlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "request handled")
}

func main() {
	var handy handlerType
	http.ListenAndServe(":8080", handy)
}
