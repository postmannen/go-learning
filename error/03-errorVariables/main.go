/*
An example that makes no sense, but to test out using error variables.
*/
package main

import "errors"
import "crypto/rand"

import "math/big"

import "log"

import "fmt"

var (
	errSmaller error = errors.New("smaller error value")
	errBigger  error = errors.New("bigger error value")
)

func someFunc() (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		log.Println("error: ", err)
	}

	if n.Int64() < 5 {
		return int(n.Int64()), errSmaller
	}

	return int(n.Int64()), errBigger

}

func main() {
	for i := 1; i <= 10; i++ {
		n, err := someFunc()
		if err != nil {
			log.Println("error: ", err)
		}
		fmt.Println("n = ", n)
		fmt.Println("------------------------------------------------")
	}
}
