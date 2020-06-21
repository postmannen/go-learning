package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/cavaliercoder/grab"
)

type server struct {
}

func newServer() *server {
	return &server{}
}

func main() {
	listenPort := flag.String("listenPort", "localhost:7777", "enter the host and port for the server to listen on")

	flag.Parse()

	s := newServer()

	http.HandleFunc("/", s.getHTTPHandler)

	if err := http.ListenAndServe(*listenPort, nil); err != nil {
		log.Printf("error: http.ListenAndServe: %v\n", err)
	}

}

// getHTTPHandler will handle the overall logic, and is the entry
// point for the whole process.
func (s *server) getHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// parse the path and file name from the request,
	u, err := url.ParseRequestURI(r.URL.String())
	if err != nil {
		fmt.Printf("error sftp: url.ParseRequestURI: %v\n", err)
	}

	remoteHost := "https://edgeos.raalabs.tech"

	// download file
	fmt.Printf("Downloading %s...\n", remoteHost+u.String())
	resp, err := grab.Get(".", remoteHost+u.String())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", r.URL.String(), err)
		return
	}

	fh, err := os.Open(resp.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: os.Open: %v\n", err)
		return
	}
	defer fh.Close()

	n, err := io.Copy(w, fh)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: io.Copy: %v\n", err)
		fmt.Fprintf(os.Stderr, "copied %n bytes\n", n)
		return
	}

	if err := os.Remove("./" + resp.Filename); err != nil {
		fmt.Fprintf(os.Stderr, "error: os.Remove: %v\n", err)
		return
	}

	fmt.Printf("Successfully downloaded to %s\n", resp.Filename)

}
