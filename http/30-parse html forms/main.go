package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", billCreateWeb)
	http.ListenAndServe(":8080", nil)
}

func billCreateWeb(w http.ResponseWriter, r *http.Request) {
	tplData := `<!DOCTYPE html>
					<form>
    				    <input type="text" name="tekstboks" value="tekstverdi">
    				    <input type="submit" name="knapp" value="verdi">
    				</form>
				`

	tpl, err := template.New("page").Parse(tplData)
	if err != nil {
		log.Printf("error: 1. failed parsing template : %v\n", err)
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("error: 2. failed executing template : %v\n", err)
	}

	r.ParseForm()
	//r.FormValue("value_of_name=XXXX>") to get the value of name=XXX. Returns only a single value
	//r.Form["<value_of_name=XXXX>"]	to get the value of name=XXX . Can get multiple values and stores them in a slice
	minInput := r.Form["knapp"]
	//log.Printf("billCreateWeb : New line button pressed, %v\n", minInput[0])

	//Check if a string of values are parsed from form. Will return runtime error if not checked since the slice is not created.
	if minInput != nil && minInput[0] == "verdi" {
		fmt.Fprintf(w, "%v\n", r.Form["tekstboks"])
	}
}
