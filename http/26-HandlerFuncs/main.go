package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type server struct {
	addr   string
	router *mux.Router //server type will use gorilla mux
}

//routes contain all the routes for the server
func (s *server) routes() {
	s.router.HandleFunc("/login", s.login())
}

func newServer() *server {
	return &server{
		addr:   ":3000",
		router: mux.NewRouter()}
}

//by making a function returning a HandlerFunc we get more flexibility compared to
//the normal 'func someName(w http.ResponseWriter, r *http.Request' way.
//The HandlerFunction's are also methods of the server object allowing us to use
//the data available trough the server struct.
//We can now also pass in arguments to the function when calling it without voilating
//the interface specification of how a HandlerFunc shall look like, since it now returns
//a HandlerFunc, and is not specified as a HandlerFunc.
func (s *server) login() http.HandlerFunc {
	var tpl *template.Template
	var err error
	var init sync.Once

	//make sure the templates are only loaded once.
	init.Do(func() {
		tpl, err = template.ParseFiles("login.html")
		if err != nil {
			fmt.Println("failed parsing template", err)
		}
	})

	return func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "login", nil)
		if err != nil {
			//testing giving errors back to the client
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}
	}
}

func main() {
	srv1 := newServer()
	srv1.router = mux.NewRouter()
	srv1.routes()
	//srv1.router.HandleFunc("/login", srv1.login())
	http.ListenAndServe(srv1.addr, srv1.router)
}
