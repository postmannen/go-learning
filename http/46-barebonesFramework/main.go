package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("Barebones")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error: ListenAndServe failed: ", err)
	}
}
