package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	certDir := flag.String("certDir", "", "specify the full path of where to store the key and certificate")
	domain := flag.String("domain", "", "the domain name to create a certificate for")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user! Your config: %+v", r.TLS)
	})

	dirPath := path.Join(*certDir, *domain)
	_, err := os.Stat(dirPath)
	if err != nil {
		err := os.MkdirAll(dirPath, 0700)
		if err != nil {
			log.Printf("error: os.MkdirAll: %v\n", err)
			return
		}
	}

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(*domain),
		Cache:      autocert.DirCache(dirPath),
	}

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go func() {
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()

	err = server.ListenAndServeTLS("", "")
	if err != nil {
		log.Printf("error: ListenAndServe: %v\n", err)
	}
}
