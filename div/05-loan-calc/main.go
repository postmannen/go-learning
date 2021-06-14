package main

import (
	"fmt"
)

type kost struct {
	mnd        int
	restSum    float64
	avdragKost float64
	renteKost  float64
}

func main() {
	//set variablene til ønskede verdier
	rest := 165000.0
	rente := 3.0
	avdrag := 3000.0

	bil := make([]kost, 1, 1)
	bil[0].restSum = rest
	fmt.Printf("Måned = %-.2v, Rest = %-.2f, avdragKost = %-.2f, renteKost = %-.2f \n", bil[0].mnd, bil[0].restSum, bil[0].avdragKost, bil[0].renteKost)

	i := 1
	for {
		tmp := kost{}

		if bil[i-1].restSum >= avdrag {
			tmp.renteKost = bil[i-1].restSum / 100 * rente / 12
			tmp.avdragKost = avdrag - tmp.renteKost
			tmp.restSum = bil[i-1].restSum - tmp.avdragKost
			tmp.mnd = i
		} else {
			tmp.renteKost = bil[i-1].restSum / 100 * rente / 12
			tmp.avdragKost = bil[i-1].restSum
			tmp.restSum = bil[i-1].restSum - tmp.avdragKost
			tmp.mnd = i
		}

		if tmp.renteKost >= avdrag {
			fmt.Println("----------------------------------------------------------------------------------------")
			fmt.Println("Du har satt totale innbetalinger til en lavere sum enn de månedlige rentekostnadene.")
			fmt.Println("Månedlige rentekostnader = ", tmp.renteKost)
			break
		}

		bil = append(bil, tmp)

		fmt.Printf("Måned = %-10.2v, Rest = %-10.2f, avdragKost = %-10.2f, renteKost = %-10.2f \n", tmp.mnd, tmp.restSum, tmp.avdragKost, tmp.renteKost)

		if tmp.restSum <= 0 {
			fmt.Println("--------------------------------------------------------------------------")
			fmt.Println("Antall måneder = ", i)
			fmt.Println("Antall år = ", float64(float64(i)/float64(12)))
			break
		}

		i++

	}

	var sumRente float64
	var sumAvdrag float64
	for _, v := range bil {
		sumRente = sumRente + v.renteKost
		sumAvdrag = sumAvdrag + v.avdragKost
	}
	fmt.Printf("Total rentekost = %-.2f\n", sumRente)
	fmt.Printf("Total sum innbetalt = %-.2f\n\n", sumAvdrag)

}
