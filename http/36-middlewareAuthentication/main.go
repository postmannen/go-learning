/*
	Testing wrapping an authenticatin Handler around a normal HandlerFunc.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

//authHandler takes a handler func to wrap with authentication, a userName and a password as input.
// Since a HandleFunc takes a HandlerFunc as it's last argument, we can create a wrapper function
// like this with a different signature which returns a function with the signature of a HandlerFunc.
// The wrapped function 'mainHandler' will be called/executed after the signature returned from
//authHandler is returned to the HandleFunc in main.
func authHandler(hf http.HandlerFunc, u string, pw string) http.HandlerFunc {
	//If authentication fails we return only a failed message, and not the main page.
	if u != "me" || pw != "pw" {
		return func(w http.ResponseWriter, r *http.Request) {
			//Writing a response code have to be done before the actual write happens, cause the
			// buffers gets flushed upon write.
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Authentication failed !!")
		}

	}
	//If authentication is ok, we return a the page requested to the HandleFunc
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Authentication OK.")
		//The hf(w,r) function will be called after the surrounding HandlerFunc is returned
		// to the HandleFunc in main.
		hf(w, r)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the main page.")
}

func main() {
	http.HandleFunc("/", authHandler(mainHandler, "mu", "pw"))
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Println("error: failed to start http server: ", err)
	}
}
