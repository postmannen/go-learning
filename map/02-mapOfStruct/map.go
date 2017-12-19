/*
map3 := map[int]myStruct1{}
map3[0] = myStruct1{firstname: "Hetti", lastname: "Duck"}
map3[1] = myStruct1{"Letti", "Duck"}
fmt.Println("map3 inneholder		: ", map3)
fmt.Println("map3[0] inneholder	: ", map3[0])
fmt.Println("	", map3[0].firstname, map3[0].lastname)

fmt.Println("\nLoop'er med for og range på map3, og skriver ut key og value :")
for i, v := range map3 {
	fmt.Println(i, v)
}
*/

package main

import "fmt"

type person struct {
	first  string
	last   string
	custnr int
}

func main() {

	people1 := map[int]person{}
	people1[0] = person{"Donald", "Duck", 0}
	people1[1] = person{"Dolly", "Duck", 1}

	for i, v := range people1 {
		fmt.Println("i = ", i, " and v = ", v)
	}

	fmt.Println(people1)
}
