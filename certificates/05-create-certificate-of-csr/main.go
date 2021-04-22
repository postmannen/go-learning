package main

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"golang.org/x/crypto/ssh"
)

func readCSR() {
	// Read the Certificate Request (CSR) from file
	csrBytes, err := os.ReadFile("/Users/bt/tmp/ca/ship3/test.csr")
	if err != nil {
		log.Printf("error: failed to read csr: %v\n", err)
		return
	}

	// The format of the data stored in the file are PEM,
	// and we need to decode it to be able to use it.
	// Variants of certificates are ASCII=PEM, Binary=DER.
	csrPemDecoded, _ := pem.Decode(csrBytes)

	// Parse the CSR from the pem decoded data.
	csr, err := x509.ParseCertificateRequest(csrPemDecoded.Bytes)
	if err != nil {
		log.Printf("error: failed to parse csr: %v\n", err)
		return
	}

	fmt.Printf("Public key: %v\n\n", csr.PublicKey)

	// generate a serial number
	serial, err := rand.Int(rand.Reader, (&big.Int{}).Exp(big.NewInt(2), big.NewInt(159), nil))
	if err != nil {
		log.Printf("error: failed to create serial: %v\n", err)
		return
	}

	// Make a template for the new x509 values, and fill in the values from the CSR.
	template := x509.Certificate{}
	template.DNSNames = csr.DNSNames
	template.IPAddresses = csr.IPAddresses
	template.EmailAddresses = csr.EmailAddresses
	template.Subject.CommonName = csr.Subject.CommonName

	template.SerialNumber = serial
	template.NotBefore = time.Now()
	template.NotAfter = time.Now().AddDate(10, 0, 0)
	template.IsCA = false
	template.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature | x509.KeyUsageCRLSign
	template.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageCodeSigning, x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}

	fmt.Printf("%v\n", csr.Subject)

	// Load the public ca cert
	caCertBytes, err := os.ReadFile("/Users/bt/tmp/ca/rootCA.pem")
	if err != nil {
		log.Printf("error: failed to read certificate file: %v\n", err)
		return
	}

	// Pem decode the public ca cert
	pemBlock, _ := pem.Decode(caCertBytes)

	// Parse the public ca certificate
	caCert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		log.Printf("error: failed to parse certificate: %v\n", err)
		return
	}

	// Read the Private RSA Key of the ca
	bPrivKey, err := os.ReadFile("/Users/bt/tmp/ca/rootCA.key")
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

	// Generate certificate
	clientCertBytes, err := x509.CreateCertificate(rand.Reader, &template, caCert, csr.PublicKey, privKeyRAW)
	if err != nil {
		log.Printf("error: failed to create certificate: %v\n", err)
		return
	}

	// Open a file to write the new certificate into.
	fh, err := os.OpenFile("/Users/bt/tmp/ca/ship3/test.pem", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Printf("error: failed to write pem file: %v\n", err)
		return
	}
	defer fh.Close()

	// Write the certificate into the file, and encode it in PEM format.
	pem.Encode(fh, &pem.Block{Type: "CERTIFICATE", Bytes: clientCertBytes})
	if err != nil {
		log.Printf("error: failed to pem encode certificate: %v\n", err)
		return
	}
}

func main() {
	readCSR()
}
