package main

import "fmt"

type struct1 struct {
	first string
	age   int
}

func main() {
	names := []string{"Donald", "Dolly", "Mikke", "Langbein", "Pluto", "Ole", "Dole", "Doffen"}
	map1 := map[string]struct1{}

	for i := range names {
		fmt.Println("Innhold av liste = ", names[i])
		map1[names[i]] = struct1{age: 100 + i}
	}

	for i, v := range map1 {
		fmt.Println("i = ", i, "v = ", v, ", og verdien på age trukket ut av struct = ", map1[i].age)
	}

}
