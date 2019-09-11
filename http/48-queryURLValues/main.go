package main

import (
	"fmt"
	"net/http"
)

// Example URL to use in browser
// http://localhost:8080/some?animal=apekatt&bird=spurv

func main() {
	http.HandleFunc("/some", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%v\n", r.URL.Query())

		// r.URL.Query returns a map["string"][]string
		// We can then check for the existense of a key directly,
		// and if it exists print out the []string value directly.
		v, ok := r.URL.Query()["animal"]
		if ok {
			fmt.Fprintf(w, "%v", v)
		}

	})

	http.ListenAndServe(":8080", nil)
}
