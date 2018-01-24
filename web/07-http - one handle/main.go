package main

import (
	"fmt"
	"net/http"
)

type aType string

func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apekatten klatrer")
	fmt.Fprintln(w, a)
}

func main() {
	var aHandler aType
	aHandler = "Er dette om en apekatt ? "
	http.ListenAndServe(":8080", aHandler)
}
