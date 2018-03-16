package main

import (
	"fmt"
)

type tractor struct {
	name  string
	power int
}

func main() {
	//creating a variable with a pointer to struct
	johnDeere := &tractor{
		name:  "John Deere",
		power: 165,
	}

	//with pointers to struct there is no need to dereference the value when used
	fmt.Println("Name of tractor : ", johnDeere.name)

	//try to change the name of the tractor which is a pointer to struct, but by its no-pointer method
	johnDeere.changeName("Fendt")
	fmt.Println("Name of tractor after trying to change with no-pointer receiver method : ", johnDeere.name)
	//this does not work since the method is working on a copy

	//try to change the name of the tractor which is a pointer to struct, but by its no-pointer method
	johnDeere.changeNamePointer("Fendt")
	fmt.Println("Name of tractor after trying to change with pointer receiver method : ", johnDeere.name)
	//this works since the method now works on a pointer to the struct, and not a copy

	//try to change the name of the tractor which is a pointer to struct, by its pointer function
	changeN(johnDeere, "Ford")
	fmt.Println("Name of tractor after trying to change with function : ", johnDeere.name)
	//this works since the functionworks on a pointer to the struct, and not a copy

}

func (t tractor) changeName(s string) {
	t.name = s
}

func (t *tractor) changeNamePointer(s string) {
	t.name = s
}

func changeN(t *tractor, s string) {
	t.name = s
}
