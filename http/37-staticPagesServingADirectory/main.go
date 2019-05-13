//Example will create an http server and server all files found in that directory.
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//create a file server, and serve the files found in ./
	fd := http.FileServer(http.Dir("./"))
	http.Handle("/", fd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: failed to start web server: ", err)
	}
}
