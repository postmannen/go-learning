package main

import (
	"fmt"
	"io"
)

type stringReader struct {
	pos  int
	data []byte
}

func newStringReader(s string) *stringReader {
	return &stringReader{
		pos:  0,
		data: []byte(s),
	}
}

// read will read from the stringreader a chunc up to the size
// of the input variable d, and then return how much read, and
// err == io.EOF when done.
func (s *stringReader) read(d []byte) (n int, err error) {
	if s.pos >= len(s.data) {
		return 0, io.EOF
	}

	// since the size of d is defined in func main() with a length
	// of 3, no more than 3 bytes will be read from s into d.
	n = copy(d, s.data[s.pos:])
	// add bytes read to position.
	s.pos += n

	return n, nil
}

func main() {
	const readSize = 3

	sr := newStringReader("The horse wondered why all the birds where flying in the same direction.")
	var horseData []byte

	for {
		d := make([]byte, readSize)

		// read into d and up to the size of d, from sr
		n, err := sr.read(d)
		if n == 0 && err == io.EOF {
			break
		}
		fmt.Println("iterating ..... bytes read = ", n, ", err = ", err, ",data read = ", d)

		// append the data read into horseData
		horseData = append(horseData, d...)

	}

	fmt.Printf("horseData contains : %v\n", string(horseData))
}
