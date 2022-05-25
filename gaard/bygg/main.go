package main

import (
	"fmt"
)

func main() {
	const (
		kornPrisPrKg = 2.7
		avlingKG     = 450
		areal        = 258

		såkornPris        = 6.5
		såkornMengdePrMål = 21
		gjødselPris       = 4.2
		gjødselMengde     = 70

		transportKornPrKg = 0.15
		diesel            = 30000

		glyfosatPrMål     = 15
		ugrassmiddelPrMål = 50
		soppmiddelPrMål   = 70

		leieArbeidSprøytePrMål   = 25
		leieArbeidGjødslingPrMål = 25
		vedlikehold              = 35000
		utleiePris               = 450

		tilskuddPrMål = 400
	)

	avlingOppgjør := int(kornPrisPrKg * avlingKG * areal)
	fmt.Printf("avlingOppgjør: %v\n", avlingOppgjør)
	såkornKost := såkornPris * såkornMengdePrMål * areal
	fmt.Printf("såkornKost: %v\n", såkornKost)
	sprøyteUgrass := (ugrassmiddelPrMål + leieArbeidSprøytePrMål) * areal
	fmt.Printf("sprøyteUgrass: %v\n", sprøyteUgrass)
	sprøyteSopp := (soppmiddelPrMål + leieArbeidSprøytePrMål) * areal
	fmt.Printf("sprøyteSopp: %v\n", sprøyteSopp)
	sprøyteKveke := (glyfosatPrMål + leieArbeidSprøytePrMål) * areal
	fmt.Printf("sprøyteKveke: %v\n", sprøyteKveke)
	gjødselKost := (gjødselPris * gjødselMengde) * areal
	fmt.Printf("gjødselKost: %v\n", gjødselKost)
	kostLeieOverGjødsling := leieArbeidGjødslingPrMål * areal
	fmt.Printf("kost leie inn for overGjødsling: %v\n", kostLeieOverGjødsling)
	transportKost := areal * avlingKG * transportKornPrKg

	kornRelaterteKostnader := int(såkornKost) + sprøyteUgrass + sprøyteKveke + sprøyteSopp + int(gjødselKost) + leieArbeidGjødslingPrMål
	fmt.Printf("korn relaterte kostnader: %v\n", kornRelaterteKostnader)
	andreKostnader := int(transportKost) + diesel + vedlikehold
	fmt.Printf("andre koststnader: transport %v + diesel %v + vedlikehold %v = %v\n", transportKost, diesel, vedlikehold, andreKostnader)
	resultat := avlingOppgjør - kornRelaterteKostnader - andreKostnader
	fmt.Printf("\nResultat fra korn: avlingOppgjør(%v) - kornRelaterteKostnader(%v) - andreKostnader(%v) = %v\n", avlingOppgjør, kornRelaterteKostnader, andreKostnader, resultat)

	tilskudd := tilskuddPrMål * areal
	fmt.Printf("tilskudd: %v\n", tilskudd)
	totaltResultat := resultat + tilskudd
	fmt.Printf("resultat(%v) + tilskudd(%v) = %v\n", resultat, tilskudd, totaltResultat)
	fmt.Println("-----------------------------")

	utLeieTotal := areal * utleiePris
	fmt.Println("utleie total:", utLeieTotal)
	fmt.Println("-----------------------------")
	fmt.Printf("Differanse drive selv (%v) og leie ut inntekt (%v) : %v\n", totaltResultat, utLeieTotal, totaltResultat-utLeieTotal)
}
