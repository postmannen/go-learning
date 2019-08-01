package main

import (
	"fmt"
	"net/http"
)

// Create HandlerFunctions that will be executed, and serve
// content to our web page. A HandlerFunc is a function with
// specific signature that the http server will accept, where
// w http.ResponseWriter is the variable for which we send all
// the data we want output'ed on the webpage, and r *http.Request
// is the incomming request from a client to a server.

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dette er en test\n")
	fmt.Fprintf(w, "og her kommer litt mere tekst")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About side test")
}

func main() {
	// Map the functions URL path and the functions together to
	// so the server knows what function to call when a specific
	// path is given.
	// The name used for this is "routes".
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the webserver listening at localhost port 7000, and
	// using the default multiplexer.
	http.ListenAndServe(":7000", nil)
}
