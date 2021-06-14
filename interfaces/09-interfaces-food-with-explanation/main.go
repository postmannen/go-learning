package main

import "fmt"

//doer interfacet har typen do(string)
type doer interface {
	do(string)
}

//grain tilfredstiller doer interfacet i og med at det har en metode 'do(s string)',
//grain blir også da av typen 'doer'
type grain struct {
	name string
}

func (g grain) do(s string) {
	fmt.Printf("%v %v\n", s, g.name)
}

type vegatable struct {
	name string
}

func (v vegatable) do(s string) {
	fmt.Printf("%v %v\n", s, v.name)
}

type meat struct {
	name string
}

func (m meat) do(s string) {
	fmt.Printf("%v %v\n", s, m.name)
}

//doingWithFood tar en type doer som input. Det vil si at den kan ta grain/meat/vegetable som input siden
//alle tre har en metode 'do(string)', og dermed tilfredstiller interfacet doer
func doingWithFood(f doer) {
	//siden ett interface er input, så kan du kjøre alle metodene som det interfacet har på input variabelen
	//metoden do tar string som input.
	f.do("koker")
}

func main() {
	tomatoes := vegatable{
		name: "tomatoes",
	}

	//doingWithFood tar interface doer som input.
	//det vil si at man må gi typen som implementerer interfacet doer som input.
	//tomatoes er av typen vegetables, som har en metode do(string), og dermed tilfredstiller interfacet
	doingWithFood(tomatoes)

	//man kan også deklarere en variabel av en interface type.
	//når man setter innholdet av variabelen, så må man også få med hvilken av typene som interfacet implementerer,
	//for interfacet i seg selv har ingen variabler, bare metoder.
	//Når variabelen av interface type er gitt innhold av en implementert type, så kan man kjøre interfacets metoder direkte på variabelen.
	var beans doer
	beans = vegatable{name: "beans"}
	beans.do("growing")
}
