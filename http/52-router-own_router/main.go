package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type hello struct {
	Language string `json:"language"`
	Saying   string `json:"saying"`
}

// server holds all the parameters that are important
// for our http server.
type server struct {
	listener net.Listener
	mux      *http.ServeMux
	hellos   []hello
}

// newServer will setup and return a new http server to use,
// with a custom listener and multiplexer.
func newServer() (*server, error) {
	var s server
	var err error

	s.listener, err = net.Listen("tcp", "localhost:8080")
	if err != nil {
		return &s, fmt.Errorf("error: failed to create net.listener: %v", err)
	}

	s.mux = http.NewServeMux()

	s.hellos = []hello{
		{Language: "SE", Saying: "tjenare"},
		{Language: "NO", Saying: "heisann"},
	}

	return &s, nil
}

// mainPage prints out all the languages
func (s *server) mainPage(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(&s.hellos)
	if err != nil {
		log.Printf("error: failed to marshal hellos: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(b))
}

func (s *server) selectLanguage(w http.ResponseWriter, r *http.Request) {
	language := strings.TrimPrefix(r.URL.Path, "/")

	if language == "" {
		s.mainPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text")

	// Check if language is defined.
	for _, v := range s.hellos {
		if v.Language == language {
			w.WriteHeader(200)
			w.Write([]byte(language))
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Country abbrevation not found"))

}

func main() {
	s, err := newServer()
	if err != nil {
		log.Printf("%v\n", err)
	}

	//s.mux.HandleFunc("/all", s.mainPage)
	s.mux.HandleFunc("/", s.selectLanguage)

	http.Serve(s.listener, s.mux)

}
