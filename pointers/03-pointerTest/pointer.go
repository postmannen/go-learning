//Dette funker ikke
package main

import "fmt"

type struct1 struct {
	struct1var1 string
}

type struct2 struct {
	*struct1
}

func main() {

	var1 := struct1{}
	var1.struct1var1 = "Donald"

	//var2.struct1var1 = "Dolly"

	fmt.Printf("%T\n", var1)

}
