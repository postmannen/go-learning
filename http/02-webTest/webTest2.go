package main

import "fmt"
import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>En header</h1>
		<p>og en linje </p>
		<p>og en linje til </p>
	`)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":7000", nil)
}
