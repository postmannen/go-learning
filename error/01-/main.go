package main

import "fmt"

func main() {
	sum, err := divide(10.0, 0.0)
	if err != nil {
		fmt.Println("Error with division : ", err)
	}
	fmt.Println("Sum = ", sum)

}

func divide(a, b float64) (sum float64, err error) {
	if b == 0.0 {
		err = fmt.Errorf("Division with Zero")
	} else {
		sum = a / b
	}

	return sum, err
}
