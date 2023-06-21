package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
	"github.com/pingcap/errors"
	"golang.org/x/crypto/ssh"
)

func main() {
	keyPath := flag.String("keyPath", "/Users/bt/.ssh/id_ed25519", "the full path to the private key file")
	flag.Parse()

	fh, err := os.Open(*keyPath)
	if err != nil {
		log.Fatalf("error: failed to open key file: %v", err)
	}

	keyBytes, err := io.ReadAll(fh)
	if err != nil {
		log.Fatalf("error: failed to read key file: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(keyBytes)
	if err != nil {
		log.Fatalf("error: failed to ssh parse private key: %v", err)
	}

	// nats.SignatureHandler
	sig := func(bytes []byte) ([]byte, error) {
		sig, err := signer.Sign(rand.Reader, bytes)
		if err != nil {
			return nil, err
		}
		return sig.Blob, err
	}

	uJWTHandler := func() (string, error) {
		return "apekatt", nil
	}

	opt := nats.UserJWT(uJWTHandler, sig)
	fmt.Printf("opt: %#v\n", opt)

	nConn, err := nats.Connect("localhost:4222", opt)
	if err != nil {
		//log.Fatalf("error: nats.Connect failed: %v", err)
	}

	fmt.Printf("nConn: %#v\n", nConn)

	nkeys.CreateUser()
}

// ---------------------------------------------------

func connectNats() error {
	type natsConfig struct {
	}

	var natsOpts []nats.Option

	fh, _ := os.Open("ssh key file")

	// use the host ssh key instead
	b, err := io.ReadAll(fh)
	if err != nil {
		return fmt.Errorf("failed to read host key file:%v", err)
	}

	// parse the private key
	signer, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return err
	}

	// add a custom signer to the NATS connection options
	natsOpts = append(natsOpts, nats.UserJWT(
		func() (string, error) {
			return "myJWT", nil
		}, func(bytes []byte) ([]byte, error) {
			sig, err := signer.Sign(rand.Reader, bytes)
			if err != nil {
				return nil, err
			}
			return sig.Blob, err
		}))

	conn, err := nats.Connect("localhost:4222", natsOpts...)
	if err != nil {
		return errors.Annotate(err, "failed to connect to NATS")
	}

	_, err = conn.JetStream()
	if err != nil {
		return fmt.Errorf("failed to create a jet stream context: %v", err)
	}

	//a.conn = conn
	//a.js = js

	return nil
}

// PublicKeyForSigner is used to convert the ssh public key to a User NKey for use in generating JWT tokens
func PublicKeyForSigner(signer ssh.Signer) (string, error) {
	key := signer.PublicKey()

	marshalled := key.Marshal()
	seed := marshalled[len(marshalled)-32:]

	encoded, err := nkeys.Encode(nkeys.PrefixByteUser, seed)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}
