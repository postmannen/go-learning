package main

import "fmt"

type inhabitant interface {
	//if you have the same method as an interface, you are also of that interfaces type
	//so..both 'swede' and 'norwegian' are of type 'inhabitant' since they have the method 'speak'
	speak()
}

type norwegian struct {
	language string
}

func (n norwegian) speak() {
	fmt.Println("Norwegians speak ", n.language)
}

type swede struct {
	language string
}

func (s swede) speak() {
	fmt.Println("Swedes speak ", s.language)
}

func speakAnyLanguage(i inhabitant) {
	fmt.Println("speakAnyLanguage sier : ", i)
}

func main() {
	myNorwegian := norwegian{language: "Norwegian"}
	mySwede := swede{language: "Swedish"}

	fmt.Println(mySwede, myNorwegian)

	//both myNorwegian and mySwede are of types that have the method 'speak',
	//and therefore they can use the speak method of the interface,
	//since they are also of the interface type since they have the same method
	speakAnyLanguage(myNorwegian)
	speakAnyLanguage(mySwede)

}
