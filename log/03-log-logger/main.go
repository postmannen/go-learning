package main

import (
	"log"
	"os"
)

type hello struct {
	l *log.Logger
}

func newHello(l *log.Logger) *hello {
	h := hello{
		l: l,
	}

	return &h
}

func main() {
	logger := log.New(os.Stderr, "test-logger ", log.Ldate|log.Ltime|log.Lshortfile)
	h := newHello(logger)

	h.l.Printf("some error: %v\n", "the error message")

}
