package main

import (
	"fmt"
)

/*
The vehicleReader interface
*/
type vehicleReader interface {
	read([]byte) (int, error)
}

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
 pre declared slice of bytes with a predefined size as input....but its empty. Since a slice is just a pointer to an
 underlying array with cap and len we are actually passing something back via the methods input! Since a slice
 is just a pointer to an array, and in this case an array defined within main, the changes to the slice
 in the method - will happen directly in the slice variable in main.
 The return part of the method is only used to tell how well the reading went. The return of data is happening via
 the methods input variable.
*/
func (tractor) read(b []byte) (n int, err error) {
	engineController := "Tractor : The engine is running"
	copy(b, engineController)
	return len(engineController), nil
}

type car struct {
	name string
}

func (car) read(b []byte) (n int, err error) {
	engineController := "Car : The engine is running"
	copy(b, engineController)
	return len(b), nil
}

func main() {
	//Create a buffer to hold the information read from the tractors controller
	buf := make([]byte, 100)

	//Create and read the tractor controller by calling its read method. The reading will update buf variable here in main since its a slice.
	johnDeere := tractor{}
	johnDeere.read(buf)
	fmt.Println(string(buf))

	//create and read the car controller by calling its read method. The reading will update buf variable here in main since its a slice.
	bmw := car{}
	bmw.read(buf)
	fmt.Println(string(buf))

	/*
		Since both tractor and car have a read method, and theyre read methods have the same specifications as
		 the read method in the 'vehicleReader' interface, they become a vehicleReader as well as beeing tractor's and car's.
		Then we should be able to pass both a tractor and a car to a function that takes a type vehicleReader as its input.
	*/
	readVehicleController(johnDeere)
}

/*
The readVehicle controller will read both car's and tractor's since theyr'e both fullfulling the requirement to become a vehicleReader.
*/

func readVehicleController(v vehicleReader) {
	buf := make([]byte, 100)
	v.read(buf)
	fmt.Println("Reading the controller of the vehicle : ", string(buf))
}
