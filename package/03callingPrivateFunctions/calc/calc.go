package calc

import "fmt"

// Calc which is an exported function will call the unexported
// calc function to do the actual calculation which will also
// deliver the result.
func Calc(a int, b int) int {
	return calc(a, b)
}

func calc(a int, b int) int {
	fmt.Println("Calling private function calc to do the sum")
	return a + b
}
