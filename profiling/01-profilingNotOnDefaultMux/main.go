package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
}

func main() {
	// To avoid exposing the profiling information out to the internet
	// we can start sharing the profiling information with it's own
	// ListenAndServe running on it's own port, and we give it the
	// pprof.Index which is a handler as the HandlerFunc to serve.
	go func() {
		log.Println(
			http.ListenAndServe("localhost:6060", http.HandlerFunc(pprof.Index)),
		)
	}()
	http.HandleFunc("/test", myHandler)
	http.ListenAndServe("localhost:8000", nil)
}
