package main

import "fmt"

type walker interface {
	walk(miles int)
}

type camel struct {
	Name string
}

type horse struct {
	Name string
}

func (c camel) walk(miles int) {
	fmt.Println(c.Name, "is walking with a bag on his back", miles, "miles")
}

func (h horse) walk(miles int) {
	fmt.Println(h.Name, "is walking with a person on his back", miles, "miles")
}

func (h horse) jump(height int) {
	fmt.Println(h.Name, "is jumping ", height)
}

func (h horse) speak() {
	fmt.Println("The horse wants to say something")
}

func longWalk(w walker) {
	w.walk(500)
	w.walk(500)
}

func main() {
	camel1 := camel{"Bill"}
	longWalk(camel1)

	horse1 := horse{"Blakken"}
	horse1.speak()
	longWalk(horse1)
	//Longwalk som er en method av type walker interface, kan ta i mot både type
	//horse og camel som input. Både horse og camel er også av type walker i og
	//med at de også er av type walker via methoden speak som de begge har.

	//interface tilater å ta i mot en value av flere forskjellige typer.

	//*******************Litt testing under her*********************

	var horse2 walker
	horse2 = horse{"Lynet"}
	//'horse2' er av typen 'walker' som er ett interface.
	//Interfacet 'walker' har metoden 'walk', som også typen 'horse' har.
	//derfor kan vi lage en 'horse' med interfacet walker, og gi den ett 'name'
	//siden struct 'horse' har 'name' som en variabel.
	horse2.walk(300)
	//'horse2' kan benytte methoden walk fra 'horse', men kan ikke bruke metoden
	//jump, da jump ikke er definert på struct camel, og kan derfor heller ikke
	//legges inn på interfacet walker. Hadde camel også hatt jump, så kunne main
	//lagt inn jump på interfacet walker.
	fmt.Printf("horse2 er av typen %T\n", horse2)
	//som man ser av utskriften, så er horse2 av typen main.horse, og ikke av
	//typen walker

	horse3 := horse{"Skyggen"}
	fmt.Printf("horse3 som er laget av typen horse har fått typen %T\n", horse3)
	horse3.jump(253)
}
