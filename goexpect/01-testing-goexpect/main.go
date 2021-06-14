package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"time"

	expect "github.com/google/goexpect"

	"github.com/google/goterm/term"
)

const (
	timeout = 10 * time.Minute
)

var (
	addr = flag.String("address", "10.0.0.7", "address of telnet server")
	pass = flag.String("pass", "bbllhdnu", "password to use")
	cmd  = flag.String("cmd", "sh run", "command to run")

	passRE   = regexp.MustCompile("Password:")
	promptRE = regexp.MustCompile("%")
)

func main() {
	flag.Parse()
	fmt.Println(term.Bluef("Telnet 1 example"))

	e, _, err := expect.Spawn(fmt.Sprintf("telnet %s", *addr), -1)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	fmt.Println("telnet session started")

	a, b, err := e.Expect(passRE, time.Second*5)
	fmt.Println("a=", a, "b=", b)
	if err != nil {
		log.Println("error: ", err)
	}
	err = e.Send(*pass + "\n")
	if err != nil {
		log.Println("error: ", err)
	}
	a, b, err = e.Expect(promptRE, timeout)
	fmt.Println("a=", a, "b=", b)
	if err != nil {
		log.Println("error: ", err)
	}
	err = e.Send(*cmd + "\n")
	if err != nil {
		log.Println("error: ", err)
	}
	result, _, _ := e.Expect(promptRE, timeout)
	err = e.Send("exit\n")
	if err != nil {
		log.Println("error: ", err)
	}

	fmt.Println(term.Greenf("%s: result: %s\n", *cmd, result))
}

