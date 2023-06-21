package lexlang

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// ---------------------------------------------------------------------
// lexer
// ---------------------------------------------------------------------

type lexer struct {
	scanner       *bufio.Scanner
	indexPosition int // indexPosition for where the position we're working in the dataSource slice.
	line          string
	err           error
}

// newAnimal will take a []string with animal names as input,
// and return a pointer to an animalParser struct.
func newLexer(r io.Reader) *lexer {
	s := bufio.NewScanner(r)

	return &lexer{
		scanner:       s,
		indexPosition: 0,
	}

}

func (a *lexer) start() {
	// We need to kickstart the process by reading the first value
	// from the input slice.
	fn := a.readNextLine()
	for {
		// execute the current function, and put the return value
		// into fn, so we can exexute that on the next iteration.
		fn = fn()

		// If no function was returned, and we received the value
		// <nil> we know that we are done, and can return to main.
		if fn == nil {
			return
		}
	}
}

type lexFunc func() lexFunc

func (a *lexer) readNextLine() lexFunc {
	if a.scanner.Scan() {
		a.line = a.scanner.Text()
		return a.printLine
	}

	if err := a.scanner.Err(); err != nil {
		fmt.Printf("%v\n", err)
		return a.error
	}

	return a.allDone
}

func (a *lexer) error() lexFunc {
	fmt.Printf("%v\n", a.err)
	return nil
}

func (a *lexer) printLine() lexFunc {
	fmt.Printf("%v\n", a.line)
	return a.readNextLine
}

func (a *lexer) allDone() lexFunc {
	return nil
}

// ---------------------------------------------------------------------
// Line lexer
// ---------------------------------------------------------------------

type lineLexer struct {
	indexPosition int
	currentChr    byte
	reader        io.Reader
	line          string
	err           error
}

func newLineLexer(s string) *lineLexer {
	r := strings.NewReader(s)

	return &lineLexer{
		reader:        r,
		indexPosition: 0,
	}

}

func (l *lineLexer) readNextChr() lexFunc {
	b := make([]byte, 1)

	_, err := l.reader.Read(b)
	if err != io.EOF {
		return l.done()
	}
}

func (l *lineLexer) done() lexFunc {
	return nil
}
