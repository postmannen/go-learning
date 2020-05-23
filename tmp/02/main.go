package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type exitCode int

const (
	EXIT_SUCCESS         exitCode = 1
	EXIT_FAILURE         exitCode = 2
	EXIT_INVALIDARGUMENT exitCode = 3
	EXIT_NOTIMPLEMENTED  exitCode = 4
	EXIT_NOPERMISSION    exitCode = 5
	EXIT_NOTINSTALLED    exitCode = 6
	EXIT_NOTCONFIGURED   exitCode = 7
	EXIT_NOTRUNNING      exitCode = 8
)

var exitCodeStrings = map[exitCode]string{
	0: "EXIT_SUCCESS",
	1: "EXIT_FAILURE",
	2: "EXIT_INVALIDARGUMENT",
	3: "EXIT_NOTIMPLEMENTED",
	4: "EXIT_NOPERMISSION",
	5: "EXIT_NOTINSTALLED",
	6: "EXIT_NOTCONFIGURED",
	7: "EXIT_NOTRUNNING",
}

// getExitCode takes the nr. at the end of the error, and
// returns the string value found in the map for that error
// code.
func getExitCode(e error) string {
	s := strings.Split(e.Error(), " ")
	sLast := s[len(s)-1]
	c, err := strconv.Atoi(sLast)
	if err != nil {
		log.Println("**** Failed to convert string to int")
	}

	code := exitCode(c)
	return exitCodeStrings[code]
}

func startWG() error {
	cmd := exec.Command("systemctl", "start", "wg0.service")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func statusWG() error {
	cmd := exec.Command("systemctl", "status", "wg0.service")

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := statusWG()
	if err != nil {
		log.Println("main : statusWG failed : ", err)
		fmt.Println("THE EXIT CODE WAS : ", getExitCode(err))

		err := startWG()
		if err != nil {
			log.Println("main: startWG failed : ", err)
		} else {
			log.Println("main: startWG OK : ", err)
		}
	} else {
		log.Println("main: wireguard already running")
	}

}
