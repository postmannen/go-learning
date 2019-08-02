package main

import "fmt"
import "net/http"

type aType string

// Any type that have a method with the following signature....
// ServeHTTP(w http.ResponseWriter, r *http.Request)
// is fullfilling the Handler interface, and can be used to
// serve a webpage.

func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Utskrift fra ServeHTTP funksjon for aType")
}

type bType string

func (b bType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Utskrift fra ServeHTTP funksjon for bType")

}

func main() {
	var aVar aType = "Monkey"
	http.Handle("/", aVar)

	var bVar bType = "Badger"
	http.Handle("/b", bVar)
	http.ListenAndServe(":8080", nil)

}
