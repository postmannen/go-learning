package main

import (
	"fmt"

	"github.com/postmannen/training/web/04-metod-handler-2nd/web"

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
