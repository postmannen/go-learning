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

	ch := make(chan string)

	go readBlock(br, ch)

	for {
		select {
		case s, ok := <-ch:
			fmt.Println("--- OUTPUT:", s)
			if !ok {
				ch = nil
			}
		}
		if ch == nil {
			break
		}
	}

}

//readBlock will check for ending brackets,and if no such
// exist on the same line it will combine the line with
// the next one, and returned the combined result.
func readBlock(r *bufio.Reader, ch chan string) {
	fmt.Println("--------------------------------------------------------------------")
	var currentLine string
	var combinedLine string
	for {
		l, _, err := r.ReadLine()
		if err != nil {
			log.Println("Error: failed reading line: ", err)
			break
		}
		nextLine := string(l)
		nextLine = strings.TrimSpace(nextLine)

		startOK := strings.HasPrefix(currentLine, "<")
		simpleEndOK := strings.HasSuffix(currentLine, ">")
		peekStartOK := strings.HasPrefix(string(nextLine), "<")
		_ = fmt.Sprintf("=== StartOK:%v, simpleEndOK:%v ,peekStartOK:%v\n", startOK, simpleEndOK, peekStartOK)

		if (startOK && !simpleEndOK) || !startOK {
			combinedLine = fmt.Sprintf("%v %v", combinedLine, currentLine)
		} else {

		}

		//If the line got both start and end brackets, send it to the channel
		ch <- currentLine

		//The reader allways lies one step ahead of where we work, all work
		// is now done for the current line, so swap it with the next, re run
		// for loop, and read a new line into next line.
		currentLine = nextLine
	}
	fmt.Println("--------------------------------------------------------------------")
	close(ch)
}
