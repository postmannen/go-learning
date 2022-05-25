package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ed25519"
)

func main() {

	public, private, _ := ed25519.GenerateKey(nil)

	b64Public := base64.RawStdEncoding.EncodeToString(public)
	b64Private := base64.RawStdEncoding.EncodeToString(private)

	fmt.Printf("public key: %v\n", b64Public)
	fmt.Printf("private key: %v\n", b64Private)

	args := []string{"bash", "-c", "uname -a"}
	argsString := strings.Join(args, " ")
	message := []byte(argsString)

	sig := ed25519.Sign(private, message)
	if !ed25519.Verify(public, message, sig) {
		log.Printf("valid signature rejected")
	}

	s := base64.RawStdEncoding.EncodeToString(sig)
	fmt.Printf("signature : %v\n", s)

	sigFromB64, err := base64.RawStdEncoding.DecodeString(s)
	if err != nil {
		log.Printf(" * failed to decode bas64: %v\n", err)
		return
	}

	if ed25519.Verify(public, message, sigFromB64) {
		log.Printf("signature of message accepted")
	}

	wrongMessage := []byte("wrong message")
	if ed25519.Verify(public, wrongMessage, sigFromB64) {
		log.Printf("signature of different message accepted")
	}
}
