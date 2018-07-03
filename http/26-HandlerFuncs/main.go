package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type server struct {
	addr   string
	router *mux.Router //server type will use gorilla mux
	user   []user
}

//routes contain all the routes for the server
func (s *server) routes() {
	s.router.HandleFunc("/login", s.login())
	s.router.HandleFunc("/register", s.register())
}

func newServer() *server {
	return &server{
		addr:   ":3000",
		router: mux.NewRouter()}
}

//By making a function returning a HandlerFunc we get more flexibility compared to
//the normal 'func someName(w http.ResponseWriter, r *http.Request' way.
//The HandlerFunction's are also methods of the server object allowing us to use
//the data available trough the server struct.
//We can now also pass in arguments to the function when calling it without voilating
//the interface specification of how a HandlerFunc shall look like, since it now returns
//a HandlerFunc, and is not specified as a HandlerFunc.
//
//The login method
func (s *server) login() http.HandlerFunc {
	var tpl *template.Template
	var err error
	var init sync.Once

	//make sure the templates are only loaded once.
	init.Do(func() {
		log.Println("Loading the template from filesystem.")
		tpl, err = template.ParseFiles("login.html")
		if err != nil {
			fmt.Println("failed parsing template", err)
		}
	})

	log.Println("Before initializing the HandlerFunction")
	return func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "login", nil)
		if err != nil {
			//testing giving errors back to the client
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}
		log.Println("After initialized the HandlerFunction")

		u, err := readUserLoginForm(r)
		if err != nil {
			log.Println("Could not readUserLoginForm", err)
		}
		fmt.Println(u)

		err = s.checkUserExist(u)
		if err != nil {
			fmt.Fprintf(w, "The user %v does not exist !", u.Email)
		} else {
			//login user things should come here !!!
		}
	}
}

func (s *server) register() http.HandlerFunc {
	return nil
}

func (s *server) checkUserExist(u user) (err error) {
	//check if user exist
	if len(s.user) != 0 {
		for _, v := range s.user {
			if v.Email == u.Email {
				break
			}
		}
	} else {
		return errors.New("Could not find user")
	}
	return nil
}

func readUserLoginForm(r *http.Request) (u user, err error) {
	var decoder = schema.NewDecoder()
	err = r.ParseForm()
	if err != nil {
		log.Println("error: Parsing form: ", err)
		return u, err
	}
	err = decoder.Decode(&u, r.PostForm)
	if err != nil {
		log.Println("Error Decoding form : ", err)
		return u, err
	}
	return u, nil
}

//User to hold all user info
type user struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Submit   string `schema:"submit"`
	Cancel   string `schema:"cancel"`
	loggedIn bool
}

func main() {
	srv1 := newServer()
	srv1.router = mux.NewRouter()
	srv1.routes()
	//srv1.router.HandleFunc("/login", srv1.login())
	http.ListenAndServe(srv1.addr, srv1.router)
}
