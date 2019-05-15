package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://vg.no" + r.RequestURI)
	if err != nil {
		log.Println("error: get: ", err)
	}
	defer resp.Body.Close()

	//The header is a map, we take all the pieces from the received header, and put the same
	// key's and values into our own header.
	for name, values := range resp.Header {
		w.Header()[name] = values
	}

	w.WriteHeader(resp.StatusCode)

	//test to print some on top of the received body.
	fmt.Fprint(w, "apekatt")

	n, err := io.Copy(w, resp.Body)
	fmt.Println("********Copied bytes = ", n)
	if err != nil {
		log.Println("error: copy: ", err)
	}
}

func main() {
	http.HandleFunc("/", pageHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("listen and serve : ", err)
	}

}
