// Idea for creating a buffer.
// Not done.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
)

var errClosed = errors.New("buffer is closed")

type buffer struct {
	bb     *bytes.Buffer
	closed bool
}

func NewBuffer(buf []byte) *buffer {
	b := buffer{
		bb: bytes.NewBuffer(buf),
	}
	return &b
}

func (b *buffer) Read(p []byte) (int, error) {
	if b.closed {
		return 0, errClosed
	}
	n, err := b.bb.Read(p)

	return n, err
}

func (b *buffer) Write(p []byte) (int, error) {
	if b.closed {
		return 0, errClosed
	}
	n, err := b.bb.Write(p)

	return n, err
}

// Todo
func (b *buffer) Close() error {
	b.closed = true

	return nil
}

// ************* main **************

func main() {
	var bb []byte
	buff := NewBuffer(bb)

	{
		_, err := buff.Write([]byte{'a'})
		if err != nil {
			log.Fatal(err)
		}
	}
	{
		_, err := buff.Write([]byte{'b'})
		if err != nil {
			log.Fatal(err)
		}
	}

	{
		b := make([]byte, 1)
		buff.Read(b)
		fmt.Println("read: ", string(b))
	}
	buff.Close()
	{
		b := make([]byte, 1)
		n, err := buff.Read(b)
		if err == errClosed {
			fmt.Printf("n: %v, err: %v, error type: %T\n", n, err, err)
		}
		fmt.Println("read: ", string(b))
	}

}
