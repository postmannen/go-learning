package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type server struct {
	tempFileFolder string
	sftpRootPath   string
	sftpUser       string
	sftpIP         string
	sftpPass       string
	sftpAuth       string
}

func newServer(tempFileFolder string, sftpRootPath string, sftpUser string, sftpIP string, sftpPass string, sftpAuth string) *server {
	return &server{
		tempFileFolder: tempFileFolder,
		sftpRootPath:   sftpRootPath,
		sftpUser:       sftpUser,
		sftpIP:         sftpIP,
		sftpPass:       sftpPass,
		sftpAuth:       sftpAuth,
	}
}

func main() {
	version := "0.13"

	listenPort := flag.String("listenPort", "localhost:7777", "enter the host and port for the server to listen on")
	tempFileFolder := flag.String("tempFileFolder", "./", "If needed, specify where to store tmp files. All files are automatically deleted")
	sftpRootPath := flag.String("sftpRootPath", "/www", "specify the directory that will be served as the root path on the sftp server")
	sftpUser := flag.String("sftpUser", "webreporeader", "ssh/sftp server username")
	sftpPass := flag.String("sftpPass", "", "password to authenticate to sftp server, will default to key if no password is given")
	sftpIP := flag.String("sftpIP", "51.120.77.187", "ip address or hostname of the sftp server")
	showVersion := flag.Bool("version", true, "Show the current version")
	sftpAuth := flag.String("sftpAuth", "password", "choose between password or pem, defaults to password")

	flag.Parse()

	if *showVersion {
		versionPrinter(version)
	}

	s := newServer(*tempFileFolder, *sftpRootPath, *sftpUser, *sftpIP, *sftpPass, *sftpAuth)

	http.HandleFunc("/", s.getHTTPHandler)

	if err := http.ListenAndServe(*listenPort, nil); err != nil {
		log.Printf("error: http.ListenAndServe: %v\n", err)
	}

}

func versionPrinter(v string) {
	fmt.Printf("Version %v\n", v)

}

// getHTTPHandler will handle the overall logic, and is the entry
// point for the whole process.
func (s *server) getHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// parse the path and file name from the request,
	u, err := url.ParseRequestURI(r.URL.String())
	if err != nil {
		fmt.Printf("error sftp: url.ParseRequestURI: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// initiate the sftp session, and download the file.
	remoteFileWithPath := s.sftpRootPath + u.String()
	sshClient, err := s.sshRequest()
	if err != nil {
		fmt.Println("error sftp: sshRequest: ", err)
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer sshClient.Close()

	// create new SFTP client
	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		fmt.Printf("error sftp: sftp.NewClient: %v\n", err)
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}

	remoteFile, err := sftpClient.Open(remoteFileWithPath)
	if err != nil {
		fmt.Printf("error sftp: sshClient.Open path = %v : %v\n", remoteFileWithPath, err)
		http.Error(w, err.Error(), http.StatusGatewayTimeout)
		return
	}
	defer remoteFile.Close()

	n, err := io.Copy(w, remoteFile)
	if err != nil {
		fmt.Printf("error sftp: io.Copy(w, client) failed after %v bytes: %v", n, err)
		http.Error(w, err.Error(), http.StatusNoContent)
	}
	fmt.Printf("sftp-info: bytes copied = %v, file = %v\n", n, remoteFile)
}

func (s *server) sshRequest() (*ssh.Client, error) {
	port := ":22"

	var signer ssh.Signer

	if s.sftpAuth != "password" {
		var err error
		signer, err = getSigner()
		if err != nil {
			return nil, fmt.Errorf("error: getSigner %v", err)
		}
	}

	config, err := s.newSSHConfig(s.sftpAuth, signer)
	if err != nil {
		return nil, fmt.Errorf("error: newSSHConfig: %v", err)
	}

	// connect
	conn, err := ssh.Dial("tcp", s.sftpIP+port, config)
	if err != nil {
		log.Printf("error: ssh.Dial: %v", err)
		return nil, err
	}

	return conn, nil
}

func (s *server) newSSHConfig(sftpAuth string, signer ssh.Signer) (*ssh.ClientConfig, error) {
	if sftpAuth == "password" {
		return &ssh.ClientConfig{
			User: s.sftpUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(s.sftpPass),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}, nil
	}

	if sftpAuth == "key" {
		return &ssh.ClientConfig{
			User: s.sftpUser,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(signer),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}, nil
	}

	return nil, fmt.Errorf("error: newSSHConfig: sftpAuth flag values password or key missing")
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
