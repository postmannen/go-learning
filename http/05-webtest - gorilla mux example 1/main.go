package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//SayHelloWorld test
func SayHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

//Greet test
func Greet(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("Hello %s !", name)))
}

//ProcessPathVariables test
func ProcessPathVariables(w http.ResponseWriter, r *http.Request) {

	// break down the variables for easier assignment
	vars := mux.Vars(r)
	name := vars["name"]
	job := vars["job"]
	age := vars["age"]
	w.Write([]byte(fmt.Sprintf("Name is %s ", name)))
	w.Write([]byte(fmt.Sprintf("Job is %s ", job)))
	w.Write([]byte(fmt.Sprintf("Age is %s ", age)))
}

func main() {
	mx := mux.NewRouter()

	mx.HandleFunc("/", SayHelloWorld)
	mx.HandleFunc("/{name}", Greet)

	//to handle URL like
	//http://website:8080/person/Boo/CEO/199

	//http://website:8080/person/Boo/CEO/199 <- if age > 199, will cause 404 error

	mx.HandleFunc("/person/{name}/{job}/{age:[0-199]+}", ProcessPathVariables)

	http.ListenAndServe(":8080", mx)
}
