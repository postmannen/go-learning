package main

import (
	"fmt"
)

type celcius float64

//String method satisfies the fmt.Stringer Interface.
//It takes no arguments and returns a string,
//and is used for controlling the format of a variable. For example how many decimals etc.
func (c celcius) String() string {
	return fmt.Sprintf("%.2f C", c)
}

func main() {
	c := celcius(10.0)
	d := string("20")
	fmt.Println(c)
	fmt.Println(d)

}
