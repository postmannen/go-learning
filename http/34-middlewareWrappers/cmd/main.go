package main

import (
	"net/http"

	"github.com/postmannen/go-learning/http/34-middlewareWrappers/pkg/rootpage"
	"github.com/postmannen/go-learning/http/34-middlewareWrappers/pkg/wrapper"
)

func newServer(addr string) *http.Server {
	return &http.Server{
		Addr: addr,
	}
}

func main() {
	srv := newServer(":8080")

	http.HandleFunc("/", rootpage.RootHandler)
	http.HandleFunc("/wrapper", wrapper.MyWrapper(rootpage.RootHandler))

	srv.ListenAndServe()

}
