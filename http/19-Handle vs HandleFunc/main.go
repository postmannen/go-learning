package main

import (
	"net/http"
)

type myStruct struct {
	myField string
}

func (m *myStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(m.myField))
}

func (m myStruct) myHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(m.myField))
}

func main() {
	//Here we are starting a web page using http.Handle, and the ServeHTTP method of the struct
	http.Handle("/b", &myStruct{myField: "Made using http.Handle, and passing the address of the struct with input field as handler"})

	//Here we are starting a web page using http.HandleFunc to servere another Handler method on myStruct
	myVar := myStruct{myField: "Made using http.HandleFunc, and passing a method of a predeclared struct as Handler"}
	http.HandleFunc("/a", myVar.myHandleFunc)

	http.ListenAndServe(":8080", nil)
}
