package main

import (
	"fmt"
	"net/http"
)

func myFunc(isAdmin bool) http.HandlerFunc {
	if !isAdmin {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Not logged in as admin")
		}
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Logged in as admin")
	}
}

func main() {
	http.HandleFunc("/", myFunc(false))
	http.ListenAndServe(":3000", nil)
}
