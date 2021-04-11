package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	value1 := []byte("This is the first value of some random data")
	shaValue1 := sha256.Sum256(value1)
	fmt.Printf("%v\n", shaValue1)

	b64Value1 := base64.StdEncoding.EncodeToString(shaValue1[:])
	fmt.Printf("%d\n", []byte(b64Value1))
}
