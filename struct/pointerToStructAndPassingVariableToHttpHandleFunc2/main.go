package main

import (
	"fmt"
	"net/http"
)

type name struct {
	firstName string
	lastName  string
}

func changeMe(p *name) {
	p.firstName = "Scrooge"
}

func main() {
	p1 := &name{
		firstName: "Donald",
		lastName:  "Duck",
	}

	p2 := name{
		firstName: "Dolly",
		lastName:  "Duck",
	}

	changeMe(p1)
	fmt.Println("After changeme the firsName = ", p1.firstName)

	server := http.Server{}
	server.Addr = "127.0.0.1:8080"

	http.HandleFunc("/1", p1.page1)
	server.ListenAndServe()

	http.HandleFunc("/2", p2.page2)

}

func (n *name) page1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Content of 'n' = ", n)
	fmt.Fprintf(w, "firstName = %v, lastName = %v", n.firstName, n.lastName)
}

func (n name) page2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Content of 'n' = ", n)
	fmt.Fprintf(w, "firstName = %v, lastName = %v", n.firstName, n.lastName)
}
