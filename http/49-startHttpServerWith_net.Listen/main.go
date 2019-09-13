/*
Testing out net.Listen and http.Serve,
and making a web server that will run for 10 seconds,
and then shut down.
*/
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "The main page.\n")
}

func main() {
	http.HandleFunc("/", mainHandler)

	// we create a net.Listen type to use later with the http.Serve function.
	nl, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("error: starting net.Listen: ", err)
	}

	// start the web server with http.Serve instead of the usual http.ListenAndServe
	go http.Serve(nl, nil)

	// select will block here (unless there had been a default case),
	// and it will wait for 10 seconds before closing down nl,
	// and the web server is close down.
	select {
	case <-time.Tick(time.Second * 10):
		fmt.Println("closing down web server")
		nl.Close()
	}
}
