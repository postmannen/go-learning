package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func handlerFunction1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "written by function1")
}

func main() {

	//HandlerFunc is a type with the signature "type Handlerfunc func(http.ResponeWriter, r* http.Request)"
	var varHandlerFunc1 http.HandlerFunc = handlerFunction1

	//Signature : http.Handle(pattern string, handler http.Handler)
	//A http.Handlerfunc type is a type of type function, which also have ServeHTTP method.
	//Since it has a ServeHTTP method it is also of type 'handler', and can be passed to http.Handle
	http.Handle("/Handle_myHandlerFunc1", varHandlerFunc1)

	//Signature : http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	//A http.HandleFunc will take any function that has input arguments of type http.ResponeWriter and *http.Request
	http.HandleFunc("/HandleFunc_function1", handlerFunction1)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
