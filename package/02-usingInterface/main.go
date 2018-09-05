/*
The whole purpose if this program is to limit what variables and methods gets
exported out of the package.
Since the printerType is not exported its methods will not be exported either.

We then create an interface containing printerType's methods.
The interface APrinter is exported, and will be available outside the package.
Since we use the New() function which returns a not exported printerType via
the APrinter interface type, we are able to export a private type via the exported
interface type.
All methods belonging to the interface type APrinter will be availabe to the new
variable create in main.
*/
package main

import (
	"fmt"

	"github.com/postmannen/go-learning/package/02-usingInterface/printsome"
)

func main() {
	ap := printsome.New()
	fmt.Printf("ap is of type %T \n", ap)
	ap.PrintA()
	ap.PrintB()

}
