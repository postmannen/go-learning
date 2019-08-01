package main

import "fmt"
import "net/http"

// Create a HandlerFunc, and use fmt.Fprintf to directly write
// some content to our webpage.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>A header</h1>
		<p>A line </p>
		<p>and another line </p>
	`)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":7000", nil)
}
