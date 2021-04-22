package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("/Users/bt/tmp/ca/ship3/ship3.csr")
	if err != nil {
		log.Printf("error: failed to read csr: %v\n", err)
		return
	}

	// Variants of certificates are ASCII=PEM, Binary=DER.
	p, rest := pem.Decode(b)
	fmt.Printf("rest: %v\n\n", rest)
	fmt.Printf("p.Type: %v\n\n", p.Type)
	fmt.Printf("p.Headers: %v\n\n", p.Headers)

	// Get the CSR from the pem decoded data.
	csr, err := x509.ParseCertificateRequest(p.Bytes)
	if err != nil {
		log.Printf("error: failed to parse csr: %v\n", err)
		return
	}

	fmt.Printf("csr.Subject: %#v\n\n", csr.Subject)
	// fmt.Printf("*csr: %+v\n\n", csr)

}
