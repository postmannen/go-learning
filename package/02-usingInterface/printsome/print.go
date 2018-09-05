package printsome

import (
	"fmt"
)

//APrinter is an interface to do some testing
type APrinter interface {
	PrintA()
	PrintB()
}

//New creates a new printerType
func New() APrinter {
	return printerType{}
}

type printerType struct {
}

//PrintA printsA
func (printerType) PrintA() {
	fmt.Println("Printing from PrintA")
}

//PrintB printsB
func (printerType) PrintB() {
	fmt.Println("Printing from PrintB")
}
