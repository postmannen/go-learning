package main

import (
	"fmt"
	"log"
	"net/http"
)

func logme(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		println("inlogme", r.URL)

		log.Println("Before")
		defer log.Println("After")
		h.ServeHTTP(w, r)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	println("in mainHandler", r.Header)
	fmt.Fprintln(w, "ape")
}

func main() {
	http.HandleFunc("/", logme(mainHandler))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error: ListenAndServe: ", err)
	}
}
