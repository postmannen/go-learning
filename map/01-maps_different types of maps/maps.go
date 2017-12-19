package main

import "fmt"

type myStruct1 struct {
	firstname string
	lastname  string
}

var map1 map[string]int //declare the map

func main() {
	defer fmt.Println("") //for å få en ekstra linkeskift i bunnen
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("map1 not initialized	: ", map1) //not initialized map. Conaints NIL
	map1 := make(map[string]int)
	fmt.Println("map1 initialized	: ", map1)
	fmt.Println("-------------------------------------------------------------------")

	map2 := map[string]int{} //declaring and initializing a map in one go
	fmt.Println("map2	: ", map2)
	map2["Donald"] = 30 //adding key and value to map
	fmt.Println("map2	: ", map2)
	fmt.Println("-------------------------------------------------------------------")

	//Map with int as key, and struct1 as value
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
	fmt.Println("-------------------------------------------------------------------")

	//map med key av int, og pointer til struct som value
	map4 := map[int]*myStruct1{}
	fmt.Println("map4 som har int som key, og peker til myStruct1 inneholder : ", map4)
	fmt.Println()
	fmt.Println(`Legger til : map4[0] = &myStruct1{"Netti", "Duck"}`)

	map4[0] = &myStruct1{firstname: "Netti", lastname: "Duck"}
	fmt.Println("map4 inneholder da : ", map4)

	fmt.Printf("\n")
	fmt.Println(`Legger til : map4[1] = &myStruct1{"Letti", "Duck"}`)

	map4[1] = &myStruct1{"Letti", "Duck"}
	//map4 skal da inneholde 2 index verdier (0 og 1), med en peker til struct adresse i hver

	fmt.Println("map4 inneholder da : ", map4)

	fmt.Println()
	fmt.Println("Utskrift av  :   map4[0].firstname, map4[0].lastname")
	fmt.Println("\t\t og map4[1].firstname, map4[1].lastname")
	fmt.Println("\t\t\t", map4[0].firstname, map4[0].lastname)
	fmt.Println("\t\t\t", map4[1].firstname, map4[1].lastname)
	fmt.Println("-------------------------------------------------------------------")

	fmt.Println("Map of maps")
	fmt.Println("Example : someMap := make(map[string]map[string]int)")

	//land-fylke-tettsted
	map5 := map[string]map[string]string{}
	//map5[1] = map[string]string{}
	//map5[1]["Østold"] = "spydeberg"

	//Man kan sjekke om key har verdi nil, hvis ikke initialiser map med key.
	/* 	if m["uid"] == nil {
	    		m["uid"] = map[string]T{}
			}
	*/

	someKey := "Norge"
	//Sjekk om det er en verdi i første key, og hvis det ikke er det så legg inn verdi der, og initialiser map
	if val, ok := map5[someKey]; ok {
		fmt.Println("The key did exist")
	} else {
		fmt.Println("The key did not exist, and the value of the key was ", val)
		map5[someKey] = map[string]string{}
	}

	map5["Norge"]["Østfold"] = "Spydeberg"
	fmt.Println(map5)
	fmt.Println("-------------------------------------------------------------------")

	//map of slice
	map6 := map[string]int{}
	map6["apekattObservert"] = []int{}

	fmt.Println(map6)

}
