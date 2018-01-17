package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"runtime"
	"time"
)

type pageVariables struct {
	Date string
	Time string
	MyOS string
}

type people struct {
	Name    string
	Surname string
	Address string
	County  string
	Mail    string
	Phone   string
}

func user(w http.ResponseWriter, r *http.Request) {

}

func homepage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()              // find the time right now
	HomePageVars := pageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
		MyOS: runtime.GOOS,
	}

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	//func ParseFiles(filenames ...string) (*Template, error)
	//ParseFiles parses the named files and associates the resulting
	// templates with t. If an error occurs, parsing stops and the returned
	// template is nil; otherwise it is t. There must be at least one file.
	if err != nil { // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	//func (t *Template) Execute(wr io.Writer, data interface{}) error.
	//Execute applies a parsed template to the specified data object,
	// writing the output to wr. If an error occurs executing the template or
	//writing its output, execution stops
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func mainMenu() {
	fmt.Printf("1. Register user \n")
	fmt.Print("Choose menu nr. : ")

}

func main() {
	/*http.HandleFunc("/", homepage)
	http.HandleFunc("/user", user)
	log.Fatal(http.ListenAndServe(":8080", nil))*/

	mainMenu()

}
