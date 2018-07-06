package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type server struct {
	id int
}

//MyHandler testing
func (s *server) MyHandler(w http.ResponseWriter, r *http.Request) {
	s.id++

	// Get a session. Get() always returns a session, even if empty.
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[s.id] = s.id
	// Save it before we write to the response/return from the handler.
	session.Save(r, w)

	fmt.Fprintln(w, s.id)
}

func main() {
	router := mux.NewRouter()
	s := &server{id: 0}

	router.HandleFunc("/", s.MyHandler)
	http.ListenAndServe(":3000", router)

}
