/*
Create a HandleFunc which returns a type HandlerFunc.
This lets us wrap more things together before executing the actual handler.
It also gives more closure around the code written.
*/

package main

import (
	"fmt"
	"net/http"
	"sync"
	"text/template"
)

const aTemplate string = `
{{define "someTemp"}}
 <h1>This is written from aTemplate</h1>
{{end}}
{{template "someTemp" .}}
`

//server is the type to hold all information about the server
type server struct {
	name     string
	once     sync.Once
	hostPort string
	tmpl     *template.Template
}

func (s *server) handleSome() http.HandlerFunc {
	s.once.Do(func() {
		fmt.Println("This is only going to be written once !!")

		tmp, err := template.New("myTemplate").Parse(aTemplate)
		if err != nil {
			fmt.Println("The template loading or parsing failed", err)
			panic(err)
		}
		s.tmpl = tmp

	})

	//if the loading of the template above was ok, then return the HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is just a test of %v", s.name)
		s.tmpl.Execute(w, nil)

	}
}

func main() {
	myServer := server{
		name:     "My personal server",
		hostPort: ":3000",
	}

	http.HandleFunc("/some", myServer.handleSome())
	http.ListenAndServe(myServer.hostPort, nil)
}
