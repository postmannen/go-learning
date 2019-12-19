// The gimmick here is that, unique among go's arithmetic-ish operators,
// << does not require that its parameters have the same type. It rather
//requires that the right operand is an unsigned integer type.
//
//It can't require that they be the same type, because it has to be unsigned,
// and you have to be able to shift ints.
// So 1 << uint64(x) is an expression shifting a constant left by x,
// but nothing in the expression constrains the type of the constant.
// A constant which does not have its type constrained and could be an int is an int.
// if you had this in a func returning (res uint64), and did res = (1 << ui) - 1,
// you'd get a uint64.
package main

import (
	"fmt"
)

func main() {
	var ui uint64
	fmt.Printf("%v, %T\n", ui, ui)
	res := 0 << ui
	fmt.Printf("%v, %T\n", res, res)
}
