package main

import (
	"fmt"
)

func skrivTall(tall int) (utTall int) {
	fmt.Println("Tallet fra funksjon = ", tall)
	utTall = tall + 100
	return
}

func main() {
	fmt.Println("----------------------------------------")
	var mittTall int
	mittTall = 99

	fmt.Println("Tallet før funksjonen er kjørt = ", mittTall)

	mittTall = skrivTall(mittTall)

	fmt.Println("Tallet etter at funksjonen er kjørt er nå = ", mittTall)

}
