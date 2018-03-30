package main

import (
	"fmt"
)

func skrivTall(tall *int) {
	fmt.Println("Tallet fra funksjon = ", *tall)
	*tall = *tall + 100
}

func main() {
	fmt.Println("----------------------------------------")
	var mittTall int
	mittTall = 99

	fmt.Println("Tallet før funksjonen er kjørt = ", mittTall)

	skrivTall(&mittTall)

	fmt.Println("Tallet etter at funksjonen er kjørt er nå = ", mittTall)

}
