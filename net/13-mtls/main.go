package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func serverStart() error {
	cert, err := tls.LoadX509KeyPair("./certs/server-cert.pem", "./certs/server-key.pem")
	if err != nil {
		return fmt.Errorf("error: failed to open cert: %v", err)
	}

	certPool := x509.NewCertPool()
	pemCABytes, err := ioutil.ReadFile("./certs/ca-cert.pem")
	if err != nil {
		return fmt.Errorf("error: failed to read ca cert: %v", err)
	}

	if !certPool.AppendCertsFromPEM(pemCABytes) {
		return fmt.Errorf("error: failed to append ca to cert pool")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientCAs:    certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	ln, err := tls.Listen("tcp", "127.0.0.1:43000", config)
	if err != nil {
		return fmt.Errorf("error: failed to start server listener: %v", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return fmt.Errorf("error: conn.Accept failed: %v", err)
		}

		for {
			b := make([]byte, 1500)
			_, err := conn.Read(b)
			if err != nil {
				return fmt.Errorf("error: conn.Read failed: %v", err)
			}

			fmt.Printf(" * read: %s\n", b)
		}
	}

}

func clientStart() error {
	cert, err := tls.LoadX509KeyPair("./certs/client-cert.pem", "./certs/client-key.pem")
	if err != nil {
		return fmt.Errorf("error: failed to open cert: %v", err)
	}

	certPool := x509.NewCertPool()
	pemCABytes, err := ioutil.ReadFile("./certs/ca-cert.pem")
	if err != nil {
		return fmt.Errorf("error: failed to read ca cert: %v", err)
	}

	if !certPool.AppendCertsFromPEM(pemCABytes) {
		return fmt.Errorf("error: failed to append ca to cert pool")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:43000", config)
	if err != nil {
		return fmt.Errorf("error: tls.Dial failed: %v", err)
	}
	defer conn.Close()

	b := []byte("some piece of data")

	_, err = conn.Write(b)
	if err != nil {
		return fmt.Errorf("error: conn.Write failed: %v", err)
	}

	return nil

}

func main() {
	go func() {
		err := serverStart()
		if err != nil {
			log.Printf("%v\n", err)
		}
	}()

	time.Sleep(time.Second * 1)

	err := clientStart()
	if err != nil {
		log.Printf("%v\n", err)
	}

	time.Sleep(time.Second * 1)
}
