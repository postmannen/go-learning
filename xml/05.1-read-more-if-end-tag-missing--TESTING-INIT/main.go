//Reading an xml file, and if there are multiple lines that
// belong together we add them together so lexing them later
// will be easir.
// For example a description block can span several lines
// each ending with a line break. Then we remove all line
// breaks and combine them together before we return it to
// as a single line.
//
// 05.1 Testing with preparing a new file to be used in main
// with init function
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const originalFile string = "ardrone3.xml"
const tmpFile string = "./testFile.txt"

func main() {
	//create a temporary file to work on
	prepareTmpFile()

}

//prepareTmpFile will create a tmp file with the original
// data, but all leading spaces will be removed making it
// easyer to lex it later.
// A <description> tag will also be added to the lines that
// have no tag.
// If there are more lines with no tags spanning multiple
// lines, they will be combined to a single longer line.
func prepareTmpFile() {
	fh, err := os.Open(originalFile)
	if err != nil {
		log.Println("Error: failed opening file: ", err)
	}
	defer fh.Close()

	tfh, err := os.Create(tmpFile)
	if err != nil {
		log.Println("Error: creating file: ", err)
	}
	defer tfh.Close()

	WriteTmpFile(fh, tfh)
}

//WriteTmpFile will prepare a new file without ending and
// leading spaces. Will also combine lines that belong
// together but are normally seperated with carriage return.
//
// Taking original file, and tmp file as input.
func WriteTmpFile(fh *os.File, tfh *os.File) error {
	var err error
	br := bufio.NewReader(fh)

	for {
		s, err := readBlock(br)
		if err != nil {
			log.Println("Error: ", err)
			break
		}
		fmt.Println(s)
		//_, err = bw.WriteString(s)
		_, err = tfh.WriteString(s)

		if err != nil {
			log.Fatal("Error: writing to tmp file: ", err)

		}
	}
	return err
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

		//Checks the first 4 characters of the next line.
		peek, _ := r.Peek(4)
		peekString := strings.TrimSpace(string(peek))
		lString = strings.TrimSpace(string(l))
		//fmt.Printf("== lString ==:%v\n", lString)

		startOK := strings.HasPrefix(lString, "<")
		simpleEndOK := strings.HasSuffix(lString, ">")
		peekStartOK := strings.HasPrefix(string(peekString), "<")
		_ = fmt.Sprintln(startOK, simpleEndOK, peekStartOK)

		// In the IF below just specify when you want to run once, or break out.
		// Breaking out means read no more lines this run.
		//
		if (startOK && simpleEndOK) || //check if line contains "<" and ">", indicating complete tag line, or...
			peekStartOK { //check if next line contains "<", indicating new bracket on next line
			tmpString = fmt.Sprintf("%v %v", tmpString, lString)

			// If the finnished line don't have any brackets at all, we assume it is
			// a description, so we add new tags called <description> & </description>.
			//
			tmpString = strings.TrimSpace(tmpString)
			startOK := strings.HasPrefix(tmpString, "<")
			simpleEndOK := strings.HasSuffix(tmpString, ">")
			if !startOK && !simpleEndOK {
				tmpString = fmt.Sprintf("<description>%v</description>", tmpString)
			}

			//Since strings.Trimspace removed all carriage returns, we now add
			//a new one to the line we want to return.
			tmpString = fmt.Sprintf("%v\n", tmpString)
			break
		}
		tmpString = fmt.Sprintf("%v %v", tmpString, lString)
	}
	fmt.Println("--------------------------------------------------------------------")
	return tmpString, nil
}
