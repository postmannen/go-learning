/*
The purpose of this exercise is to understand the differences between : handle, handler, handlefunc, handlerfunc.
	1. type http.Handler
	2. function http.Handle
	3. type http.HandlerFunc
	4. function http.HandleFunc
	5. using http.HandlerFunc to http.Handle
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Creating a type of empty struct which is to become a http.Handler by attaching a ServeHTTP method to it
type myHandler struct{}

//To satisfy the http.Handler interface, and become a type http.Handler a type must have a method with the
//signature : ServeHTTP(http.ResponseWriter, *http.Request
//Here we give myHandler such a method, and turn it into a type http.Handler
func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Written by myHandler's ServeHTTP method")
}

//Here we create a function that satisfies the HandlerFunc type
//Signature : type HandlerFunc func(ResponseWriter, *Request)
func handlerFunction1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "written by function1")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "The page you were looking for where not found")
}

func main() {
	//EXAMPLE 1
	//The function http.Handle takes a type http.Handler.
	//Since our myHandler type has a ServeHTTP(ResponsWriter, *Request) method it satisfies the Handler interface,
	//and becomes a Type Handler aswell.
	var varMyHandler myHandler

	//EXAMPLE 2
	//Signature : http.Handle(pattern string, handler http.Handler)
	http.Handle("/Handle_myHandler", varMyHandler)

	// EXAMPLE 3
	//HandlerFunc is a type with the signature "type Handlerfunc func(http.ResponeWriter, r* http.Request)"
	//The type Handlerfunc has a method ServeHTTP attached, so varHandlerFunc1 becomes both
	// a type http.HandlerFunc and a type http.Handler since it satisfies the http.Handler interface.
	var varHandlerFunc1 http.HandlerFunc = handlerFunction1

	//EXAMPLE 4
	//Signature : http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	//A http.HandleFunc will take any function that has input arguments of type http.ResponeWriter and *http.Request
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/HandleFunc_function1", handlerFunction1)

	//EXAMPLE 5
	//Signature : http.Handle(pattern string, handler http.Handler)
	//A http.Handlerfunc type is a type of type function, which also have ServeHTTP method.
	//Since it has a ServeHTTP method it is also of type 'handler',
	// and we can pass it into the http.Handle function.
	http.Handle("/Handle_myHandlerFunc1", varHandlerFunc1)

	//starting the first listener in its own go routine so the code doesn't block for starting the next below.
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	//--------------------With Gorilla Mux---------------------------
	//The same as over but with Gorilla Mux instead of default mux

	router := mux.NewRouter()

	//EXAMPLE 2.1
	//Signature : http.Handle(pattern string, handler http.Handler)
	router.Handle("/Handle_myHandler", varMyHandler)

	//EXAMPLE 4.1
	//Signature : http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
	//A http.HandleFunc will take any function that has input arguments of type http.ResponeWriter and *http.Request
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	router.HandleFunc("/HandleFunc_function1", handlerFunction1)

	//EXAMPLE 5.1
	//Signature : http.Handle(pattern string, handler http.Handler)
	//A http.Handlerfunc type is a type of type function, which also have ServeHTTP method.
	//Since it has a ServeHTTP method it is also of type 'handler',
	// and we can pass it into the http.Handle function.
	router.Handle("/Handle_myHandlerFunc1", varHandlerFunc1)

	//EXAMPLE 6.1
	//The function "notFound" got the correct signature to become a HandlerFunc, so we make it into one :)
	var NotFoundHF http.HandlerFunc = notFound
	//Again, since the type HandlerFunc got a ServeHTTP method that satisfies the Handler interface it also becomes a Handler,
	// and we can pass that into the mux's NotFoundHandler
	router.NotFoundHandler = NotFoundHF

	err := http.ListenAndServe(":8090", router)
	if err != nil {
		panic(err)
	}
}
