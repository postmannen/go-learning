/*
The moose loves capital M's, and will read all lowercase 'm' to uppercase 'M'.
The purpose of this program is to test io.Reader, and turning one type of reader into another type of reader.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
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

func (m moose) Read(p []byte) (n int, err error) {
	/*
		since m.reader is of type io.Reader we will use io.Reader's read method here.
		We will read one byte, and if succesful read we will replace any lowercase 'm' with an uppercase 'M'....since a moose reads like that.

		We are not using a pointer receiver here, since we will not alter the content of moose. If we had used a pointer receiver here we
		would also have to make the moose in TEST1 a pointer to moose *moose, since a method with pointer receivers only accepts type pointer. A
		non-pointer receiever accepts both pointer and non-pointer type's.
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
		mooseTranslate := make([]byte, 16)
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
		If the Read method of moose was a pointer receiver we would have to change the assignement of firstMoose to a pointer variable *moose with :
		firstMoose := &moose{
			reader: strings.NewReader("marvelous moose's moves silently more often than other mamals"),
		}
	*/
	firstMoose := moose{
		reader: strings.NewReader("marvelous moose's moves silently more often than other mamals"),
	}
	fmt.Println(mooseRead(firstMoose))

	fmt.Println("-----------------------TEST2-----------------------------------")
	/*
		Here we call the mooseRead function with a type moose as input.
		This will work since moose is also a a Reader by having its own Read method, and we can use it directly in mooseRead function.
	*/

	myMoose := newMoose(strings.NewReader("moose is mighty magic"))
	fmt.Println(mooseRead(myMoose))

	fmt.Println("-----------------------TEST3-----------------------------------")
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
		Here in TEST4 we open a file which is of type os.File, then we can turn a type os.File into a type moose with the function newMoose.
		The reason we can do this is that newMoose accepts any type that is a Reader by having a Read method, and by that satisfying
		the io.Reader interface.
		By doing this we can see that mooseRead will use moose's read method, and not os.File's read method like it did in TEST3.
	*/
	fmt.Println("-----------------------TEST4-----------------------------------")
	aFile, err := os.Open("mosesTheMoose.txt")
	if err != nil {
		fmt.Println("error: opening file = ", err)
	}
	defer aFile.Close()

	//Turn aFile of type os.File into a moose
	yetAnotherMoose := newMoose(aFile)

	println(mooseRead(yetAnotherMoose))

	fmt.Println("-----------------------TEST5-----------------------------------")
	resp, err := http.Get("http://www.dustyfeet.com/")
	if err != nil {
		fmt.Println("main: TEST5 : http.Get : ", err)
		os.Exit(1)
	}

	/*
		resp.Body is also of type io.Reader, that means we can make a moose of it
	*/
	fifthMoose := newMoose(resp.Body)
	fmt.Println(mooseRead(fifthMoose))

	fmt.Println("-----------------------TEST END-----------------------------------")

}
