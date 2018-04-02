/*
The moose loves capital M's, and will read all lowercase 'm' to uppercase 'M'.
The purpose of this program is to test io.Reader, and turning one type of reader into another type of reader.
*/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type moose struct {
	reader io.Reader
}

/*
creates a new moose.
Take the input of type reader, and return a pointer to *moose where the moose struct field 'reader' is filled with the content of the input.
*/
func newMoose(reader io.Reader) *moose {
	return &moose{reader: reader}
}

func (m *moose) Read(p []byte) (n int, err error) {
	/*
		since m.reader is of type io.Reader we will use io.Reader's read method here.
		We will read one byte, and if succesful read we will replace any lowercase 'm' with an uppercase 'M'....since a moose reads like that.
	*/
	n, err = m.reader.Read(p)
	if err != nil {
		return n, err
	}

	for i := 0; i < n; i++ {
		if p[i] == 'm' {
			p[i] = 'M'
		}
	}

	return n, nil
}

/*
mooseRead will do the complete reading of the byteSlice given to it as input, by calling the inputs underlying Read method.
*/
func mooseRead(p io.Reader) string {
	mooseTranslated := []byte{}

	for {
		mooseTranslate := make([]byte, 8)
		n, err := p.Read(mooseTranslate)
		if err != nil {
			fmt.Println("mooseRead: Read error = ", err)
			break
		}
		fmt.Printf("mooseRead: p is of type %T, and characters read = %v \n", p, n)

		mooseTranslated = append(mooseTranslated, mooseTranslate...)
	}
	fmt.Println("mooseRead: length of mooseTranslated = ", len(mooseTranslated))
	return fmt.Sprint(string(mooseTranslated))

}

func main() {
	fmt.Println("-----------------------TEST1-----------------------------------")
	/*
		Here we call the mooseRead function with a type moose as input.
		This will work since moose is also a a Reader by having its own Read method, and we can use it directly in mooseRead function.
	*/

	myMoose := newMoose(strings.NewReader("moose is mighty magic"))
	fmt.Println(mooseRead(myMoose))

	fmt.Println("-----------------------TEST2-----------------------------------")
	/*
		Since os.File is also a Reader we can pass a file of type os.File directly into the mooseRead function.
		Here we create anotherMoose, but anotherMoose will be of type *os.File
		Calling mooseRead will work since it accepts any io.Reader, but it will use os.File's Read method,
		and not moose's read method, since anotherMoose is not a moose :-)
	*/
	anotherMoose, err := os.Open("martinTheMoose.txt")
	if err != nil {
		fmt.Println("error: opening file = ", err)
	}
	defer anotherMoose.Close()

	println(mooseRead(anotherMoose))

	/*
		Here in TEST3 we open a file which is of type os.File, then we can turn a type os.File into a type moose with the function newMoose.
		The reason we can do this is that newMoose accepts any type that is a Reader by having a Read method, and by that satisfying
		the io.Reader interface.
		By doing this we can see that mooseRead will use moose's read method, and not os.File's read method like it did in TEST2.
	*/
	fmt.Println("-----------------------TEST3-----------------------------------")
	aFile, err := os.Open("mosesTheMoose.txt")
	if err != nil {
		fmt.Println("error: opening file = ", err)
	}
	defer aFile.Close()

	//Turn aFile of type os.File into a moose
	yetAnotherMoose := newMoose(aFile)

	println(mooseRead(yetAnotherMoose))
	fmt.Println("-----------------------TEST END-----------------------------------")

}
