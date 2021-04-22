package main

import (
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {
	m := &autocert.Manager{
		Cache:      autocert.DirCache("./certs"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("dev.erter.org"),
	}
	s := &http.Server{
		Addr:      ":https",
		TLSConfig: m.TLSConfig(),
	}
	s.ListenAndServeTLS("", "")
}
