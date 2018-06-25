package main

import (
	"net/http"

	"github.com/postmannen/go-learning/http/001-calhounCourse/controllers"

	"github.com/postmannen/go-learning/http/001-calhounCourse/views"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

//func signup(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	must(signupView.Render(w, nil))
//}

var homeView *views.View
var contactView *views.View
var signupView *views.View

func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")
	userC := controllers.NewUsers()
	//signupView = views.NewView("bootstrap", "views/signup.html")

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/contact", contact).Methods("GET")
	r.HandleFunc("/signup", userC.New).Methods("GET")
	r.HandleFunc("/signup", userC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
