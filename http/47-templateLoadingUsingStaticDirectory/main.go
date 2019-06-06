/*
	Test file upload on web page with MultiPart web page.
	Will read the whole file into memory,
	and write it all back to a temporary file.
*/

package main

import (
	"html/template"
	"log"
	"net/http"
)

type data struct {
	templ *template.Template
}

func newData() *data {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		log.Println("error: failed parsing template: ", err)
	}

	return &data{templ: t}
}

func (d *data) showPage(w http.ResponseWriter, r *http.Request) {
	err := d.templ.ExecuteTemplate(w, "mainHTML", nil)
	if err != nil {
		log.Println("error: executing template: ", err)
	}
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	d := newData()

	http.HandleFunc("/", d.showPage)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error: ListenAndServer failed: ", err)
	}
}
