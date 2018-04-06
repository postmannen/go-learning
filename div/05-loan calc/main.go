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
	rest := 165000.0
	rente := 4.0
	avdrag := 2500.0

	bil := make([]kost, 1, 1)
	bil[0].restSum = rest

	i := 1
	for {
		tmp := kost{}

		tmp.renteKost = bil[i-1].restSum / 100 * rente / 12
		tmp.avdragKost = avdrag - tmp.renteKost
		tmp.restSum = bil[i-1].restSum - tmp.avdragKost
		tmp.mnd = i

		bil = append(bil, tmp)

		fmt.Printf("Måned = %-.2v, Rest = %-.2f, avdragKost = %-.2f, renteKost = %-.2f \n", tmp.mnd, tmp.restSum, tmp.avdragKost, tmp.renteKost)

		if tmp.restSum <= 0 {
			fmt.Println("Antall måneder = ", i)
			fmt.Println("Antall år = ", float64(float64(i)/float64(12)))
			break
		}

		i++

	}
	//fmt.Println(bil)

}
