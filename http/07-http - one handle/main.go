package main

import (
	"fmt"
	"net/http"
)

type aType string

func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "The monkey climbs")
	fmt.Fprintln(w, a)
}

func main() {
	var aHandler aType
	aHandler = "Is this about a monkey ? "
	http.ListenAndServe(":8080", aHandler)
}
