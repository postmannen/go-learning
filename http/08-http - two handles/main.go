package main

import "fmt"
import "net/http"

type aType string

//ServeHTTP : Interfacet Handler inneholder metoden ServeHTTP(w http.ResponseWriter, r *http.Request)
//Hvis en hvilken som helst type har en slik ServeHTTP metode så tilfredstiller typen interfacet som heter Handler.
func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Utskrift fra ServeHTTP funksjon for aType")
}

type bType string

func (b bType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Utskrift fra ServeHTTP funksjon for bType")

}

func main() {
	var aVar aType = "Apekatt"
	http.Handle("/", aVar)

	var bVar bType = "Grevling"
	http.Handle("/b", bVar)
	http.ListenAndServe(":8080", nil)

}
