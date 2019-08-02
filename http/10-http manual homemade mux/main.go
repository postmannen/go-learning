package main

import (
	"fmt"
	"io"
	"net/http"
)

type aType string

func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Check what the path contains, and do a switch on it...
	switch r.URL.Path {
	case "/a":
		io.WriteString(w, "You are now at /a")
	case "/b":
		io.WriteString(w, "You are now at /b")
	}
	fmt.Fprintln(w, "\n----------------------------")
	fmt.Fprintln(w, "Path", r.URL.Path)
	fmt.Fprintln(w, "Host", r.URL.Host)
	fmt.Fprintln(w, "url", r.URL)

}

func main() {
	var myMux aType
	http.ListenAndServe(":8080", myMux)

}
