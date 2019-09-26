package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func one(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "one")
}

func two(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 1)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "two")
}

func startWebSevers() {
	// start a web server at :8080 with a / route
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", one)

		nl, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Println("error: net listen for webserver one")
		}
		http.Serve(nl, mux)

	}()

	// start a web server at :8080 with a / route,
	// which have a delay for serving the page.
	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", two)

		nl, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Println("error: net listen for webserver one")
		}
		http.Serve(nl, mux)

	}()
}

func findFastestWeb(ch chan *http.Response) {
	var found bool
	var mu sync.Mutex

	go func() {
		resp, err := http.Get("http://localhost:8080/")
		if err != nil {
			log.Println("error: http.Get for one", err)
		}

		if !found {
			mu.Lock()
			found = true
			mu.Unlock()

			ch <- resp
		}
	}()

	go func() {
		resp, err := http.Get("http://localhost:8081/")
		if err != nil {
			log.Println("error: http.Get for one", err)
		}

		if !found {
			mu.Lock()
			found = true
			mu.Unlock()

			ch <- resp
		}
	}()

}

func main() {
	startWebSevers()

	respCh := make(chan *http.Response, 1)

	findFastestWeb(respCh)

	resp := <-respCh
	defer resp.Body.Close()

	body := make([]byte, 1000)
	_, err := resp.Body.Read(body)
	if err != nil && err != io.EOF {
		log.Println("error: failed reading body", err)
	}
	fmt.Println("read body : ", string(body))
}
