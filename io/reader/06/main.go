package main

import (
	"fmt"
	"io"
	"strings"
)

type sheep struct {
	reader io.Reader
}

func (s sheep) Read(p []byte) (n int, err error) {
	fmt.Println("using sheep's read method to read ")
	n, err = s.reader.Read(p)
	//if reached end of read, or any other error occured, return bytes read, and the error
	if err != nil {
		return n, err
	}

	fmt.Println("Read : ", string(p))

	//if no error occured, that means..we are not done reading, return error set to nil
	return n, nil

}

func sheepSpeaker(p io.Reader) {
	sheepSays := []byte{}
	for {
		buf := make([]byte, 8)
		n, err := p.Read(buf)
		if err != nil {
			fmt.Println("error sheapReader : ", err)
			break
		}
		fmt.Println("sheepSpeaker : Character read = ", n)
		sheepSays = append(sheepSays, buf...)
	}
	fmt.Println("sheepSays = ", string(sheepSays))
}

func newSheep(t io.Reader) *sheep {
	return &sheep{reader: t}
}

func main() {
	//sheep1 is not a real sheep, so it will not use sheep's Read method
	sheep1 := strings.NewReader("The sheep says baaaaaa whenever it wants to")
	sheepSpeaker(sheep1)

	//sheep2 will become a sheep
	sheep2 := newSheep(strings.NewReader("This is yet another sheep"))
	sheepSpeaker(sheep2)

	//NOT WORKING !!! check
	sheep3 := newSheep(sheep2)
	sheepSpeaker(sheep3)
	fmt.Println(sheep3)
}
