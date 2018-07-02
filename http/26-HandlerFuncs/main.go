package main

import (
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
}

//routes contain all the routes for the server
func (s *server) routes() {
	s.router.HandleFunc("/login", s.login())
	s.router.HandleFunc("/loginOK", s.loginOK())
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

		var u User
		var decoder = schema.NewDecoder()
		err = r.ParseForm()
		if err != nil {
			log.Println("error: Parsing form: ", err)
		}
		err = decoder.Decode(&u, r.PostForm)
		if err != nil {
			log.Println("Error Decoding form : ", err)
		}

		fmt.Println("user info = ", u)
	}
}

func (s *server) loginOK() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "You've logged in")
	}
}

func readUserLoginForm(u *User, w http.ResponseWriter, r *http.Request) error {
	var decoder = schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		return err
	}
	err = decoder.Decode(&u, r.PostForm)
	if err != nil {
		log.Println("Error Decoding form : ", err)
	}

	fmt.Println("user info : ")
	fmt.Println("email", u.Email)
	fmt.Println("password", u.Password)
	fmt.Println("submit", u.Submit)
	fmt.Println("cancel", u.Cancel)
	return nil
}

//User to hold all user info
type User struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Submit   string `schema:"submit"`
	Cancel   string `schema:"cancel"`
}

func main() {
	srv1 := newServer()
	srv1.router = mux.NewRouter()
	srv1.routes()
	//srv1.router.HandleFunc("/login", srv1.login())
	http.ListenAndServe(srv1.addr, srv1.router)
}
