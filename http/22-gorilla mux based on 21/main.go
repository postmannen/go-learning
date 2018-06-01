package main

import (
	"fmt"
	"net/http"
)

func myHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "This is the main / ")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "this is the contact page")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "the page you requested was not found")
	}

}

func main() {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", myHandleFunc)
	http.ListenAndServe(":8080", mux)
}
