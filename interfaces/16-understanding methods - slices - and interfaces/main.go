package main

import (
	"fmt"
)

/*
This is the type describing what type of data that is actually stored.
If an interface is used with tractor type of data, then this is the underlying type that is used with its
 belonging methods.
TODO: interface will be implemented later.
*/
type tractor struct {
	name string
}

/*
This will create a read method for the tractor type.
This method could potentially be used to read any kind of information from the tractor,
 and here the info from the tractor is simulated with the string variable engineController.
The thing here that pulled me off the path to understand this, is that we are actually just passing a
 pre declared slice of bytes of a predefined size as input....but its empty. Since a slice is just a pointer to an
 underlying array with cap and len we are actually passing something back via the methods input! Since a slice
 is just a pointer to an array, and in this case an array defined within main, the changes to the slice will happen
 dirctly in the slice variable in main.
 The return part of the method is only used to tell how well the reading went. The return of data is happening via
 the methods input variable.
*/
func (tractor) read(b []byte) (n int, err error) {
	engineController := "The engine is running"
	copy(b, engineController)
	return len(engineController), nil
}

func main() {
	//Create a buffer to hold the information read from the tractors controller
	buf := make([]byte, 100)

	johnDeere := tractor{}
	//read the tractors controlle by calling its read method. The reading will update buf variable here in main since its a slice.
	johnDeere.read(buf)
	fmt.Println(string(buf))

}
