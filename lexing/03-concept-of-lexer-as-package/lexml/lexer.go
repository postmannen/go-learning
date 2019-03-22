package lexml

import (
	"errors"
	"io"
	"log"
)

//lexer defines the core of the lexer
type lexer struct {
	fileReader  []string //is the file we are reading. Each slice represents a line in the file
	lineCounter int
	lineContent string
	//someB string
}

//readLine will read the current line, and put it into l.lineContent.
// Will break out and return error if empty file given
// Will return io.EOF if end of file reached
func (l *lexer) readLine() error {
	if len(l.fileReader) == 0 {
		err := errors.New("File empty")
		return err
	}
	if l.lineCounter >= len(l.fileReader) {
		l.lineCounter = 0
		return io.EOF
	}
	l.lineContent = l.fileReader[l.lineCounter]
	l.lineCounter++

	return nil
}

//newlexer will create and return the basic structure of a lexer
func newLexer(file []string) *lexer {
	return &lexer{
		fileReader: file,
	}
}

//StartLexing should make a channel, and return it to the caller.
// Will also start up a Go routine to read the file line by line,
// and return the lines read on the channel
// The Go routine stays and does it's job reading lines, but the
// outer func goes on, exits and returns the channel to main so
// the data produced in the Go routine will be delivered to the
// the caller.
func StartLexing(file []string) chan string {
	ch := make(chan string)
	lex := newLexer(file)

	go func() {
		for {
			err := lex.readLine()
			if err != nil {
				log.Println("error: reading line: ", err)
				break
			}
			ch <- lex.lineContent
			//fmt.Println("INFO: Read from file: ", lex.lineContent)
		}
		close(ch)
	}()

	return ch
}
