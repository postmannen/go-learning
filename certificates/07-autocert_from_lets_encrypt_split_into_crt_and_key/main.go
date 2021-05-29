package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/crypto/acme/autocert"
)

func checkFileUpdated(fileRealPath string, fileUpdated chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println("Failed fsnotify.NewWatcher")
		return
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		//Give a true value to updated so it reads the file the first time.
		fileUpdated <- true
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					//testing with an update chan to get updates
					fileUpdated <- true
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(fileRealPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	daemon := flag.Bool("daemon", false, "Set to true do run in daemon mode. The certificate will be automatically renewed 30 days before it expires, and the corresponding .key and .crt file will be updated .")

	certDir := flag.String("certDir", "", "specify the full path of where to store the key and certificate")
	domain := flag.String("domain", "", "the domain name to create a certificate for")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, TLS user! Your config: %+v", r.TLS)
	})

	dirPath := path.Join(*certDir, *domain)
	_, err := os.Stat(dirPath)
	if err != nil {
		err := os.MkdirAll(dirPath, 0700)
		if err != nil {
			log.Printf("error: os.MkdirAll: %v\n", err)
			return
		}
	}

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(*domain),
		Cache:      autocert.DirCache(dirPath),
	}

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}

	go func() {
		h := certManager.HTTPHandler(nil)
		log.Fatal(http.ListenAndServe(":http", h))
	}()

	go func() {
		err = server.ListenAndServeTLS("", "")
		if err != nil {
			log.Printf("error: ListenAndServe: %v\n", err)
		}
	}()

	// The cert+key are stored in a file named by the domain.
	certRealPath := path.Join(dirPath, *domain)

	fileUpdated := make(chan bool)
	go checkFileUpdated(certRealPath, fileUpdated)

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	for {
		select {
		case <-fileUpdated:
			err := handleCertFiles(certRealPath)
			if err != nil {
				log.Printf("%v\n", err)
				os.Exit(1)
			}
		case <-sigCh:
			log.Printf("info: received signal to quit..\n")
			return
		}

		if !*daemon {
			break
		}
	}

}

func handleCertFiles(certRealPath string) error {
	// Open key+cert file for reading
	fhKeyCert, err := os.Open(certRealPath)
	if err != nil {
		return fmt.Errorf("error: failed to open cert file for reading: %v", err)
	}
	defer fhKeyCert.Close()

	// Create cert file for writing to.
	fhCert, err := os.OpenFile(certRealPath+".crt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error: failed to open key file for writing: %v", err)
	}
	defer fhCert.Close()

	// Create key file for writing to.
	fhKey, err := os.OpenFile(certRealPath+".key", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("error: failed to open key file for writing: %v", err)
	}
	defer fhKey.Close()

	scanner := bufio.NewScanner(fhKeyCert)

	for scanner.Scan() {

		// Find Key, and write it to file.
		if strings.Contains(scanner.Text(), "BEGIN EC PRIVATE KEY") {

			_, err = fhKey.WriteString(scanner.Text() + "\n")
			if err != nil {
				return fmt.Errorf("error: failed to write key file: %v", err)
			}

			for scanner.Scan() {
				_, err := fhKey.WriteString(scanner.Text() + "\n")
				if err != nil {
					return fmt.Errorf("error: failed to write key file: %v", err)
				}

				if strings.Contains(scanner.Text(), "END EC PRIVATE KEY") {
					// Advance one scanner position, for the beginning of the cert
					scanner.Scan()
					break
				}
			}
		}

		// Find certs, and write them to file.
		if strings.Contains(scanner.Text(), "BEGIN CERTIFICATE") {

			_, err = fhCert.WriteString(scanner.Text() + "\n")
			if err != nil {
				return fmt.Errorf("error: failed to write cert file: %v", err)
			}

			for scanner.Scan() {
				_, err := fhCert.WriteString(scanner.Text() + "\n")
				if err != nil {
					return fmt.Errorf("error: failed to write cert file: %v", err)
				}

				if strings.Contains(scanner.Text(), "END CERTIFICATE") {
					// Advance one scanner position, for the beginning of the cert
					// scanner.Scan()
					break
				}
			}
		}

	}

	return nil
}
