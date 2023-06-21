package lexlang

import (
	"fmt"
	"strings"
	"testing"
)

const succeed = "*** \u2713 *** "
const failed = "*** \u2717 **** "

func TestGeneral(t *testing.T) {
	// Create som input animal data
	exp1 := strings.NewReader(`
	100+200
	300+400
	`)

	// Create a new variable of type animal, and prepare it with the
	// input data.
	a := newLexer(exp1)

	// Start the state machine, and work on the input data.
	a.start()

	fmt.Printf("exp1 = %+v\n", exp1)

}
