package controllers

import (
	"fmt"
	"net/http"

	"github.com/postmannen/go-learning/http/001-calhounCourse/views"
)

//NewUsers to prepare and create a new user
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.html"),
	}
}

//Users holds all the parameter of a user
type Users struct {
	NewView *views.View
}

//New used to render the form for the create user page
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//Create is the method who actually handles the logic to create the user after
//the submit button have been pressed in the signup page (called by the New method)
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Here we are supposed to have the logic to create the user")
}
