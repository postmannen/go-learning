package main

import (
	"fmt"
	"log"
	"net/http"
)

func myHandlerFunc(userOK bool) http.HandlerFunc {
	if userOK == false {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Login failed")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Some text")
	}
}

func main() {
	var loggedIN bool
	loggedIN = false
	http.HandleFunc("/", myHandlerFunc(loggedIN))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("error: ListenAndServe", err)
	}
}
