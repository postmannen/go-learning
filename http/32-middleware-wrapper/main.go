package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct {
	address string
}

func (s *server) mainHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is the main page")
	}
}

//timeLogger is a wrapper function that takes a HandlerFunc
//as input and name it 'h', and returns a modified HandlerFunc.
//
//Since we want to wrap the original passed-in HandlerFunc with
//something more, we create a new Func to return with the signature
//of a Handlerfunc. Then we can run our time middleware code in
//the inner function, call the original function "h(w,r)", and return
//that whole as a new HandlerFunction.
func (s *server) timeLogger(h http.HandlerFunc) http.HandlerFunc {
	tStart := time.Now()

	return func(w http.ResponseWriter, r *http.Request) {
		tElapsed := time.Since(tStart)
		fmt.Fprintln(w, "Time since server started : ", tElapsed)
		h(w, r)
	}
}

func main() {
	srv := newServer(":8080")
	http.HandleFunc("/", srv.timeLogger(srv.mainHandler()))
	err := http.ListenAndServe(srv.address, nil)
	if err != nil {
		log.Println("Error: Failed starting web listener: ", err)
	}

}

func newServer(addrPort string) *server {
	return &server{
		address: addrPort,
	}
}
