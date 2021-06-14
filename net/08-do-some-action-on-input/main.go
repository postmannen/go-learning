/*
Starts a tcp listener on localhost port 9000,
for testing net.Conn, and do something based on the input
*/
package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9000")
	fmt.Println("Starting tcp listener at port 9000")
	if err != nil {
		fmt.Println("error: listener:", err)
		listener.Close()
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error: conn Accept:", err)
			conn.Close()
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		myC := string("Enter command : ")
		myCommand := []byte(myC)
		_, err := conn.Write(myCommand)
		if err != nil {
			fmt.Println("error: failed to print command to session")
		}

		//read the buffer
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error: reading Conn to buf: ", err)
			break
		}

		//convert buffer to string format
		s := string(buf)

		if s[:3] == "ape" {
			fmt.Println("du skrev ape du !")
		}

		if s[:4] == "menu" {
			conn.Write(printMenu())
		}

		//check if the 8 first letters are "hostName"
		if s[:8] == "hostname" {
			hostName, err := hostName()
			if err != nil {
				fmt.Println("error: hostname: ", err)
			}
			str := fmt.Sprint("The hostname is: ", hostName, "\n")
			conn.Write([]byte(str))
		}

		if s[:11] == "openBrowser" {
			err := openBrowser()
			if err != nil {
				fmt.Println("error: failed to open browser : ")
			}
		}

		fmt.Println("Reading ", n, "bytes, which contains the string : ", s)

	}

}

//hostName returns the hostname of the host
func hostName() (name string, err error) {
	return os.Hostname()
}

//openBrowser runs a command locally on the Os, and opens a browser
func openBrowser() (err error) {
	//os.Open("open 'https://google.com'")
	fmt.Println("Trying to open a browser")
	cmd := exec.Command("open", "https://google.com")
	err = cmd.Run()
	if err != nil {
		fmt.Println("error: cmd.Run : ")
	}
	return err
}

//printMenu prints the menu to the terminal
func printMenu() []byte {
	s := string(`	ape
	openBrowser
	hostname` + "\n")

	return []byte(s)
}
