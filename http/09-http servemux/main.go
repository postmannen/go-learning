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

	//lage en ny mux som skal ta seg av routinga av requests til riktig handler.
	mux := http.NewServeMux()
	server := http.Server{
		Addr: "127.0.0.1:8080",
		//setter man en handler på server, så vil den bare kunne hoste en handler
		//for å kunne ha flere handlere setter vi den til 'nil', som er DefaultServerMux
		//og spesifiserer handlerene utenfor.
		//Vi bruker da DefaultServeMux sin handler methode
		Handler: mux,
	}

	var myHandlerA aType
	//sette hvilken handler mux skal sende en path request til
	//med Handle så trenger man en  variabel som er av en type, som igjen har method som heter ServeHTTP med writer og request som input
	mux.Handle("/", myHandlerA)
	//med HandleFunc så trenger man ikke å ha en type med en method ServeHTTP, det holder med å ha en vanlig func
	//som tar http.ResponseWriter, og *http.Request som input
	mux.HandleFunc("/hf", anotherHandler)

	//parameterene for ListenAndServe er satt i server variabelen som er av typen http.Server som er en struct
	//Ved å opprette en struct på denne måten så kan man legge fler instillinger på ListenAndServe
	server.ListenAndServe()

}
