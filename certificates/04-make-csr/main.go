package main

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

// ** STEPS
// - Read private RSA Key
// - Parse the PEM encoded Private key into a RAW format
// - Create Certificate Request template, and fill it with values
// - Create the Certificate Request
// - Encode the Certificate Request to PEM format
// - Write the PEM encoded Certificate Request to file

func main() {
	// Read the Private RSA Key
	bPrivKey, err := os.ReadFile("/Users/bt/tmp/ca/ship3/ship3.key")
	if err != nil {
		log.Printf("error: failed to read private RSA key: %v\n", err)
		return
	}

	// Parse the PEM encoded Private key into a RAW format
	privKeyRAW, err := ssh.ParseRawPrivateKey(bPrivKey)
	if err != nil {
		log.Printf("error: failed to parse private key: %v\n", err)
		return
	}

	// Create Certificate Request template, and fill it with values
	csrTemplate := &x509.CertificateRequest{}
	// Subject, with the common name for the domain specified.
	csrTemplate.Subject = pkix.Name{CommonName: "ship3.erter.org"}
	// Since the common name can only hold one value
	csrTemplate.DNSNames = []string{"deck1.ship3.erter.org"}
	csrTemplate.IPAddresses = []net.IP{[]byte{192, 168, 0, 1}}
	csrTemplate.EmailAddresses = []string{"postmannen@gmail.com"}
	csrTemplate.Subject.CommonName = "ship3.erter.org"

	// Create the Certificate Request
	csr_asn1, err := x509.CreateCertificateRequest(rand.Reader, csrTemplate, privKeyRAW)
	if err != nil {
		log.Printf("error: failed to create csr: %v\n", err)
		return
	}

	// Encode the Certificate Request to PEM format
	// Write the PEM encoded Certificate Request to file
	fh, err := os.OpenFile("/Users/bt/tmp/ca/ship3/ship3.csr", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("error: failed to write csr file: %v\n", err)
		return
	}
	defer fh.Close()
	pem.Encode(fh, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csr_asn1})

}
