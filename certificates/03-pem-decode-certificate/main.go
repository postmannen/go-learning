package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func main() {
	b, err := os.ReadFile("/Users/bt/tmp/ca/ship1/ship1.pem")
	if err != nil {
		log.Printf("error: failed to read certificate file: %v\n", err)
		return
	}

	pemBlock, rest := pem.Decode(b)
	fmt.Printf("rest: %v\n\n", rest)
	fmt.Printf("p.Type: %v\n\n", pemBlock.Type)
	fmt.Printf("p.Headers: %v\n\n", pemBlock.Headers)

	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		log.Printf("error: failed to parse certificate: %v\n", err)
		return
	}

	fmt.Printf("cert: %#v\n\n", cert.Issuer)
}
