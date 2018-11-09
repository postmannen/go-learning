// Working but the peek function with a fixed nr is not optimal
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fh, err := os.Open("ardrone3.xml")
	if err != nil {
		log.Println("Error: failed opening file: ", err)
	}

	br := bufio.NewReader(fh)

	for {
		s, err := readBlock(br)
		if err != nil {
			log.Println("Error: ", err)
			break
		}
		fmt.Println("OUTPUT:", s, ":ENDOUTPUT")
	}

}

//readBlock will check for ending brackets,and if no such
// exist on the same line it will combine the line with
// the next one, and returned the combined result.
func readBlock(r *bufio.Reader) (string, error) {
	var lString string
	var l []byte
	var err error
	var tmpString string

	fmt.Println("--------------------------------------------------------------------")
	for {

		l, _, err = r.ReadLine()
		if err != nil {
			log.Println("Error: failed reading line: ", err)
			return "", err
		}
		fmt.Printf("===== l =====:%v\n", string(l))
		fmt.Printf("===== l =====:%v\n", l)

		peek, err := r.Peek(10)
		if err != nil {
			log.Println("Error: peek: ", err)
		}
		peekString := strings.TrimSpace(string(peek))
		lString = strings.TrimSpace(string(l))
		fmt.Printf("== lString ==:%v\n", lString)

		startOK := strings.HasPrefix(lString, "<")
		simpleEndOK := strings.HasSuffix(lString, ">")
		peekStartOK := strings.HasPrefix(string(peekString), "<")
		_ = fmt.Sprintln(startOK, simpleEndOK, peekStartOK)
		//fmt.Printf("PEEKSTARTOK:%v:%v\n ", peekStartOK, peekString)
		// If the string starts or ends with a bracket, add add that line
		// to tmpString and break out of loop. If this was a normal line
		// with a start "<" tmpString will then only contain 1 line since
		// there is no previous runs of the foor loop buffered in tmpLine.
		//
		// In the IF below just specify when you want to run once.
		//if strings.HasPrefix(lString, "<") || strings.HasSuffix(lString, "/>") {
		if (startOK && simpleEndOK) || (!startOK && peekStartOK) {
			tmpString = fmt.Sprintf("%v %v", tmpString, lString)
			fmt.Printf("===tmpString in if=====:%v\n", tmpString)
			break
		}
		tmpString = fmt.Sprintf("%v %v", tmpString, lString)
		fmt.Printf("===tmpString after if=====:%v\n", tmpString)
	}
	fmt.Println("--------------------------------------------------------------------")
	return tmpString, nil
}
