/*
The whole purpose if this program is to limit what variables and methods gets
exported out of the package.
Since the printerType is not exported its methods will not be exported either,
but we will now export them via the new APrinter interface.

We then create an interface containing printerType's methods.
The interface APrinter is exported, and will be available outside the package.
Since we use the New() function which returns a not exported printerType via
the APrinter interface type, we are able to export a private type via the exported
interface type.
All methods belonging to the interface type APrinter will be availabe to the new
variable create in main.
*/
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
