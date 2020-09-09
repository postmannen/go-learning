/*
	Testing creating a router based on the standard library.
	The overall idea here is to pick out a country code as
	the first element in the URL path and return the
	corresponding value for that country back to the user.
	If the country is not found found a status NotFound will
	be returned.
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type language struct {
	Language string `json:"language"`
	Saying   string `json:"saying"`
}

// server holds all the parameters that are important
// for our http server, so define them in it's own
// struct to keep them in one place for later reference.
type server struct {
	listener  net.Listener   // The port and address to listen on.
	mux       *http.ServeMux // Route multiplexer for the handles.
	languages []language     // The different languages defined.
}

// newServer will setup and return a new http server to use,
// with a custom listener and multiplexer.
func newServer() (*server, error) {
	var s server
	var err error

	// Create a tcp listener for the web server
	s.listener, err = net.Listen("tcp", "localhost:8080")
	if err != nil {
		return &s, fmt.Errorf("error: failed to create net.listener: %v", err)
	}

	// create a new mux for route handling
	s.mux = http.NewServeMux()

	// statically define the languages to use
	s.languages = []language{
		{Language: "SE", Saying: "tjenare"},
		{Language: "NO", Saying: "heisann"},
	}

	return &s, nil
}

// mainPage prints out all the languages defined in JSON to the user
func (s *server) mainPage(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(&s.languages)
	if err != nil {
		log.Printf("error: failed to marshal hellos: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte(b))
}

// selectLanguage pics out the language from the URL path,
// and returns the corresponding values for that language
// to the http server.
func (s *server) selectLanguage(w http.ResponseWriter, r *http.Request) {
	language := strings.TrimPrefix(r.URL.Path, "/")

	// If no language was found in the URL path
	if language == "" {
		s.mainPage(w, r)
		return
	}

	w.Header().Set("Content-Type", "text")

	// Check if language is defined.
	for _, v := range s.languages {
		if v.Language == language {
			w.WriteHeader(200)
			w.Write([]byte(language))
			return
		}
	}

	// if none of the defined languages where found in the URL path...
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
