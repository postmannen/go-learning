package main

import (
	"fmt"
	"io"
	"net/http"
)

type aType string

func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//sjekk hva url path inneholder, og kjør en switch på det
	switch r.URL.Path {
	case "/a":
		io.WriteString(w, "Du er nå på /a")
	case "/b":
		io.WriteString(w, "Du er nå på /b")
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
