package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	http.HandleFunc("/", getHTTPHandler)

	if err := http.ListenAndServe(":7777", nil); err != nil {
		log.Printf("error: http.ListenAndServe: %v\n", err)
	}

}

// getHTTPHandler will handle the http request,
// http.Response
func getHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// parse the path and file name from the request,
	u, err := url.ParseRequestURI(r.URL.String())
	if err != nil {
		log.Printf("error: url.ParseRequestURI: %v\n", err)
	}

	// initiate the sftp session, and download the file.
	remoteFileWithPath := "/home/webreporeader/www" + u.String()
	localFileWithPath := "." + u.String()
	err = sshRequest(remoteFileWithPath)
	if err != nil {
		log.Println("sshRequest: ", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// serve the file back through the response
	http.ServeFile(w, r, localFileWithPath)

	// Remove file
	f := strings.Split(localFileWithPath, "/")
	fileToRemove := f[len(f)-1]
	if err = os.Remove(fileToRemove); err != nil {
		log.Printf("error: os.Remove: %v\n", err)
	}

}

func sshRequest(fileName string) error {

	l := strings.Split(fileName, "/")
	localFile := l[len(l)-1]
	localFile = "./" + localFile
	fmt.Println("The file to be saved locally = ", localFile)

	//user := "bt"
	user := "webreporeader"
	//pass := "tester12"
	//remote := "10.0.0.110"
	remote := "51.120.77.187"
	port := ":22"

	// get host public key
	hostKey := getHostKey(remote)

	signer, err := getSigner()
	if err != nil {
		return fmt.Errorf("error: getSigner %v", err)
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	// connect
	conn, err := ssh.Dial("tcp", remote+port, config)
	if err != nil {
		log.Printf("error: ssh.Dial: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	// create new SFTP client
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Printf("error: sftp.NewClient: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	// create destination file
	dstFile, err := os.Create(localFile)
	if err != nil {
		return fmt.Errorf("error: os.Create: %v", err)

	}
	defer dstFile.Close()

	// open source file
	srcFile, err := client.Open(fileName)
	if err != nil {
		return fmt.Errorf("error: client.Open: %v", err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("error: io.Copy: %v", err)
	}
	log.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
		return fmt.Errorf("error: dstFile.Sync: %v", err)
	}

	return nil
}

func getSigner() (ssh.Signer, error) {
	fh, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"))
	if err != nil {
		log.Printf("error: failed to open pem file: %v\n", err)
		return nil, err
	}
	defer fh.Close()

	pemBytes, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Printf("error: failed to read pem file: %v\n", err)
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(pemBytes)
	if err != nil {
		log.Printf("error: failed to parse private key: %v\n", err)
		return nil, err
	}

	return signer, nil

}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}
