package main

import (
	"bufio"
	"flag"
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

type server struct {
	tempFileFolder string
	sftpRootPath   string
	sftpUser       string
	sftpIP         string
}

func newServer(tempFileFolder string, sftpRootPath string, sftpUser string, sftpIP string) *server {
	return &server{
		tempFileFolder: tempFileFolder,
		sftpRootPath:   sftpRootPath,
		sftpUser:       sftpUser,
		sftpIP:         sftpIP,
	}
}

func main() {
	listenPort := flag.String("listenPort", "localhost:7777", "enter the host and port for the server to listen on")
	tempFileFolder := flag.String("tempFileFolder", "./", "If needed, specify where to store tmp files. All files are automatically deleted")
	sftpRootPath := flag.String("sftpRootPath", "/home/webreporeader/www", "specify the directory that will be served as the root path on the sftp server")
	sftpUser := flag.String("sftpUser", "webreporeader", "ssh/sftp server username")
	sftpIP := flag.String("sftpIP", "51.120.77.187", "ip address or hostname of the sftp server")
	flag.Parse()

	s := newServer(*tempFileFolder, *sftpRootPath, *sftpUser, *sftpIP)

	http.HandleFunc("/", s.getHTTPHandler)

	if err := http.ListenAndServe(*listenPort, nil); err != nil {
		log.Printf("error: http.ListenAndServe: %v\n", err)
	}

}

// getHTTPHandler will handle the overall logic, and is the entry
// point for the whole process.
func (s *server) getHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// parse the path and file name from the request,
	u, err := url.ParseRequestURI(r.URL.String())
	if err != nil {
		log.Printf("error: url.ParseRequestURI: %v\n", err)
	}

	// Create a tmp file to store the downloaded file, also
	// imediately defer the close of the file so we clean up.
	tempFile, err := ioutil.TempFile("./", "cache-proxy.*.tmp")
	if err != nil {
		log.Printf("error: unable to create tmp file: %v\n", err)
	}
	defer os.Remove(tempFile.Name())

	fmt.Println("tempfile = ", tempFile.Name())

	// initiate the sftp session, and download the file.
	remoteFileWithPath := s.sftpRootPath + u.String()
	localFileWithPath := tempFile.Name()
	err = s.sshRequest(remoteFileWithPath, tempFile)
	if err != nil {
		log.Println("sshRequest: ", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// serve the file back through the response
	http.ServeFile(w, r, localFileWithPath)
}

func (s *server) sshRequest(fileName string, tempFile *os.File) error {
	//user := "bt"
	user := s.sftpUser
	//pass := "tester12"
	//remote := "10.0.0.110"
	remote := s.sftpIP
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

	// open source file
	srcFile, err := client.Open(fileName)
	if err != nil {
		return fmt.Errorf("error: client.Open: %v", err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(tempFile, srcFile)
	if err != nil {
		return fmt.Errorf("error: io.Copy: %v", err)
	}
	log.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = tempFile.Sync()
	if err != nil {
		return fmt.Errorf("error: dstFile.Sync: %v", err)
	}

	return nil
}

// getSigner will return a signer to use for the sftp session
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
