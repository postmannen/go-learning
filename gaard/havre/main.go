package main

import (
	"fmt"
)

func main() {
	const (
		kornPris = 2.5
		avlingKG = 500
		areal    = 258

		såkornPris    = 6.5
		gjødselPris   = 4.2
		gjødselMengde = 50

		transport = 25000
		diesel    = 30000

		leieGjødselSprøyte = 25

		vedlikehold = 25000

		utleiePris = 450
	)

	avlingOppgjør := int(kornPris * avlingKG * areal)
	fmt.Printf("avlingOppgjør: %v\n", avlingOppgjør)
	såkornKost := såkornPris * 21 * areal
	fmt.Printf("såkornKost: %v\n", såkornKost)
	sprøyteUgrass := (50 + leieGjødselSprøyte) * areal
	fmt.Printf("sprøyteUgrass: %v\n", sprøyteUgrass)
	sprøyteSopp := (70 + leieGjødselSprøyte) * areal * 0
	fmt.Printf("sprøyteSopp: %v\n", sprøyteSopp)
	sprøyteKveke := (15 + leieGjødselSprøyte) * areal
	fmt.Printf("sprøyteKveke: %v\n", sprøyteKveke)
	gjødselKost := (gjødselPris * gjødselMengde) * areal
	fmt.Printf("gjødselKost: %v\n", gjødselKost)
	overGjødsling := leieGjødselSprøyte * areal * 0
	fmt.Printf("overGjødsling: %v\n", overGjødsling)

	kornRelaterteKostnader := int(såkornKost) + sprøyteUgrass + sprøyteKveke + sprøyteSopp + int(gjødselKost) + overGjødsling
	fmt.Printf("korn relaterte kostnader: %v\n", kornRelaterteKostnader)
	andreKostnader := transport + diesel + vedlikehold
	fmt.Printf("andreKostnader: %v\n", andreKostnader)
	resultat := avlingOppgjør - kornRelaterteKostnader - andreKostnader

	fmt.Printf("\nResultat fra korn: avlingOppgjør(%v) - kornRelaterteKostnader(%v) - andreKostnader(%v) = %v\n", avlingOppgjør, kornRelaterteKostnader, andreKostnader, resultat)
	tilskudd := 400 * 260
	fmt.Printf("tilskudd: %v\n", tilskudd)
	totaltResultat := resultat + tilskudd
	fmt.Printf("resultat(%v) + tilskudd(%v) = %v\n", resultat, tilskudd, totaltResultat)
	fmt.Println("-----------------------------")
	utLeieTotal := areal * utleiePris
	fmt.Println("utleie total:", utLeieTotal)
	fmt.Println("-----------------------------")
	fmt.Printf("Differanse drive selv (%v) og leie ut inntekt (%v) : %v\n", totaltResultat, utLeieTotal, totaltResultat-utLeieTotal)

	fmt.Println("-----verdi utstyr------------")
	harv := 200000
	såmaskin := 160000
	trommel := 70000
	plog := 200000
	swift := 350000

	fmt.Println("sum redskap: ", harv+såmaskin+trommel+plog+swift)
}
