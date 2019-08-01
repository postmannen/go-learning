/*
	Using methods on a type as HandlerFunc. This will give us the
	ability to store data (state) for the web page.
*/

package main

import (
	"fmt"

	"github.com/postmannen/go-learning/web/04-methodHandler/web"

	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	myWeb := web.DataStruct{}
	myWeb.Data1 = "Apen hoppet opp i busken"

	http.HandleFunc("/a", myWeb.IndexA)
	http.HandleFunc("/b", myWeb.IndexB)
	fmt.Println(myWeb.Data1)
	server.ListenAndServe()

}
