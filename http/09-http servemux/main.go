package main

import (
	"fmt"
	"net/http"
)

type aType string

func (a aType) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dette er en utskrift aType ServeHTTP\n")
}

func anotherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dette er en utskrift fra anotherHandler som er en HandlerFunc")
}

func main() {

	// Create a new multiplexer that will take care of the routing.
	mux := http.NewServeMux()
	server := http.Server{
		Addr: "127.0.0.1:8080",
		// If only one handler is directly attacked to a http.ListenAndServe
		// then only that single Handle (page) can be served from that server.
		// To be able to serve more handles we need a multiplexer to do the
		// routing between the different Handles for us. We could use the default
		// multiplexer by specifying nil as the handler in our http.ListenAndServe
		// statement, or we could create our own

		Handler: mux,
	}

	var myHandlerA aType
	// Set the routing for the different Handlers.
	mux.Handle("/", myHandlerA)
	mux.HandleFunc("/hf", anotherHandler)

	// The parameters for the ListenAndServe are set in the http.Server struct.
	// By creating a server object like this we can put more parameters on the
	// web server.
	server.ListenAndServe()

}
