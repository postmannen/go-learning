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

//handleSome is not a real HandleFunc, it is just a normal function who returns
// a type HandlerFunc.
//The nice thing with this is that you can prepare and initialise whats needed to
// serve the web page without even calling the web page to be shown if
// something goes wrong.
//It is the returned type HandlerFunc thats beeing executed by the HandleFunc function.
func (s *server) handleSome() http.HandlerFunc {
	//If we had something that only should be loaded once for a web page we could put it in the s.once.Do
	// closure below.
	s.once.Do(func() {
		fmt.Println("This is only going to be written once !!")
	})

	//if the loading of the template above was ok, then return the HandlerFunc
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is just a test of %v", s.name)
		s.tmpl.Execute(w, nil)

	}
}

func newServer() *server {
	s := server{
		name:     "My personal server",
		hostPort: ":3000",
	}

	//putting the template loading in the newServer rather than in the handler or main, since
	// it feels like its belonging in creating a new server, and should only be loaded once.
	tmp, err := template.New("myTemplate").Parse(aTemplate)
	if err != nil {
		fmt.Println("The template loading or parsing failed", err)
		panic(err)
	}
	s.tmpl = tmp

	return &s
}

func main() {
	myServer := newServer()

	http.HandleFunc("/some", myServer.handleSome())
	http.ListenAndServe(myServer.hostPort, nil)
}
