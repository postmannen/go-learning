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

//User to hold all user info
type user struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	Submit   string `schema:"submit"`
	Cancel   string `schema:"cancel"`
	loggedIn bool
	ID       int
}

//server, holds data needed to control a http server instance
type server struct {
	addr       string
	router     *mux.Router //server type will use gorilla mux
	user       []user
	userLastID int
}

//routes contain all the routes for the server
func (s *server) routes(u *user) {
	s.router.HandleFunc("/", s.mainPage(u))
	s.router.HandleFunc("/login", s.login(u))
	s.router.HandleFunc("/register", s.register(u))
}

func newServer() *server {
	return &server{
		addr:       ":3000",
		router:     mux.NewRouter(),
		userLastID: 0,
	}
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
func (s *server) login(usr *user) http.HandlerFunc {
	var tpl *template.Template
	var err error
	var init sync.Once

	//make sure the templates are only loaded once.
	init.Do(func() {
		tpl, err = template.ParseFiles("login.html", "wrapper.html")
		if err != nil {
			fmt.Println("failed parsing template", err)
		}
	})

	return func(w http.ResponseWriter, r *http.Request) {

		err := tpl.ExecuteTemplate(w, "wrapper", usr)
		if err != nil {
			//testing giving errors back to the client
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}

		u, err := s.readUserLoginForm(r)
		if err != nil {
			log.Println("Could not readUserLoginForm", err)
		}
		fmt.Println(u)

		if found := s.checkUserExist(u); found == false {
			fmt.Fprintf(w, "The user %v does not exist !", u.Email)
		} else {
			//login user things should come here !!!
			*usr = u
			usr.loggedIn = true
			fmt.Fprintf(w, "Logged in user %v\n", usr.Email)
		}
	}
}

//Method and control of html template to handle the new user registration.
func (s *server) register(usr *user) http.HandlerFunc {
	var tpl *template.Template
	var init sync.Once
	var err error

	init.Do(func() {
		//load the template only once for a given server
		tpl, err = template.ParseFiles("register.html", "wrapper.html")
		if err != nil {
			log.Println("Error : parsing template file", err)
		}
	})

	return func(w http.ResponseWriter, r *http.Request) {
		err := tpl.ExecuteTemplate(w, "wrapper", usr)
		if err != nil {
			//testing giving errors back to the client
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Fprintln(w, err)
		}

		u, err := s.readUserLoginForm(r)
		if err != nil {
			log.Println("Error : failed reading user register form ", err)
		}

		s.user = append(s.user, user{
			Email: "nisse@nisse.com",
		})

		if found := s.checkUserExist(u); found {
			fmt.Printf("The user %v exist !", u)
			fmt.Fprintf(w, "The user %v exist !", u)
		} else {
			//if user do not exist, append the new user to the slice
			//which is the db for all the users.
			s.userLastID++
			fmt.Fprintln(w, "Could not find user appending to slice")
			s.user = append(s.user, u)
		}

	}
}

func (s *server) mainPage(usr *user) http.HandlerFunc {
	var init sync.Once
	var tpl *template.Template
	var err error
	init.Do(func() {
		tpl, err = template.ParseFiles("wrapper.html", "mainPage.html")
		if err != nil {
			log.Println("Error: parsing files for main template", err)
		}
	})

	return func(w http.ResponseWriter, r *http.Request) {
		if err := tpl.ExecuteTemplate(w, "wrapper", usr); err != nil {
			log.Println("Error: executing main template ", err)
		}
		fmt.Fprintf(w, "This is the main page\n")
	}
}

//if user exist, return true
//if user do not exist, return false
func (s *server) checkUserExist(u user) (found bool) {
	fmt.Println("The whole content of s = ", s)
	//check if user exist
	if len(s.user) != 0 { //if slice is empty..no users exists at all
		for i, v := range s.user {
			fmt.Printf("i:%v, v:%v\n", i, v)
			fmt.Printf("comparing v.Email:%v with u.Email:%v\n", v.Email, u.Email)
			if v.Email == u.Email {
				fmt.Println("--found user--")
				return true
			}
		}
	} else {
		fmt.Println("s.[]user was empty, no users to check")
	}
	return false
}

//Reads the form, and uses gorilla/schema to decode the values based
//on the user struct.
func (s *server) readUserLoginForm(r *http.Request) (u user, err error) {
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

func newUser() *user {
	return &user{
		loggedIn: false,
	}
}

func main() {
	u := newUser()

	srv1 := newServer()
	srv1.router = mux.NewRouter()
	srv1.routes(u)
	//srv1.router.HandleFunc("/login", srv1.login())
	http.ListenAndServe(srv1.addr, srv1.router)
}
