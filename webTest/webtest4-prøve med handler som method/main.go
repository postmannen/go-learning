package main

import (
	"fmt"
	"net/http"
)

type handlerType string

func (h handlerType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "request handled")
}

func main() {
	var handy handlerType
	http.ListenAndServe(":8080", handy)
}
