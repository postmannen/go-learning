package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("/Users/bt/tmp/ca/ship1/ship1.csr")
	if err != nil {
		log.Printf("error: failed to read csr: %v\n", err)
		return
	}

	p, rest := pem.Decode(b)
	fmt.Printf("rest: %v\n\n", rest)
	fmt.Printf("p.Type: %v\n\n", p.Type)
	fmt.Printf("p.Headers: %v\n\n", p.Headers)

	csr, err := x509.ParseCertificateRequest(p.Bytes)
	if err != nil {
		log.Printf("error: failed to parse csr: %v\n", err)
		return
	}

	fmt.Printf("csr.Subject: %+v\n\n", csr.Subject)
	fmt.Printf("*csr: %+v\n\n", *csr)

}
