package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

type server struct {
	addr   string
	router http.Handler
}

func login() http.HandlerFunc {
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

func newServer() *server {
	return &server{addr: ":3000", router: http.DefaultServeMux}
}

func main() {
	server1 := newServer()

	http.HandleFunc("/login", login())
	http.ListenAndServe(server1.addr, server1.router)
}
