package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type myStruct struct {
	Str1 string
	Str2 string
}

func main() {
	myData := myStruct{
		Str1: "str1 data",
		Str2: "str2 data",
	}

	//testing handlefunc with anonymous function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(os.Stdout, "and the *request looked like this", r)

		tmpl, err := template.ParseFiles("test.html")
		if err != nil {
			fmt.Println("error: parsefile : ", err)
		}
		tmpl.Execute(w, myData)

	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error: ListenAndServer: ", err)
	}
}
