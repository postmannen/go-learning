package main

import (
	"log"
	"strings"

	"golang.org/x/crypto/ed25519"
)

func main() {

	public, private, _ := ed25519.GenerateKey(nil)

	args := []string{"bash", "-c", "uname -a"}
	argsString := strings.Join(args, " ")
	message := []byte(argsString)

	sig := ed25519.Sign(private, message)

	if !ed25519.Verify(public, message, sig) {
		log.Printf("valid signature rejected")
	}

	if ed25519.Verify(public, message, sig) {
		log.Printf("signature of message accepted")
	}

	wrongMessage := []byte("wrong message")
	if ed25519.Verify(public, wrongMessage, sig) {
		log.Printf("signature of different message accepted")
	}
}
