package main

import (
	"log"

	"golang.org/x/crypto/ed25519"
)

type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func main() {
	var zero zeroReader
	public, private, _ := ed25519.GenerateKey(zero)

	message := []byte("test message")
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
