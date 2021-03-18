package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	ok, err := checkAddrInPrefix("127.0.0.100", "127.0.0.0", 24)
	if err != nil {
		log.Printf("checkAddrInPrefix failed: %v\n", err)
		return
	}

	fmt.Printf("Result = %v\n", ok)
}

// Will take an address, prefix, mask bits as it's input, and return
// true if the addr where within the specified prefix.
func checkAddrInPrefix(addr string, prefix string, maskBits int) (bool, error) {
	a, err := convertDotStringToUint32(addr)
	if err != nil {
		return false, err
	}

	p, err := convertDotStringToUint32(prefix)
	if err != nil {
		return false, err
	}

	m := convertMaskbitToUint32(maskBits)

	result := p&m == a&m

	return result, nil
}

// Convert for example a 24 bits mask to its uint32 representation
func convertMaskbitToUint32(bitsSet int) uint32 {
	restBits := 32 - bitsSet
	u := uint32(math.Pow(2, float64(bitsSet)) - 1)
	u = u << uint32(restBits)

	return u
}

// Will convert string x.x.x.x of ip address into uint32
func convertDotStringToUint32(s string) (uint32, error) {
	nSplit := make([]byte, 4)

	sSplit := strings.Split(s, ".")
	for _, v := range sSplit {
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, fmt.Errorf("error: strconv.Atoi failed: %v", err)
		}

		nSplit = append(nSplit, byte(n))
	}

	// Create an uint32 value of the ip address.
	var addrUint32 uint32
	for _, v := range nSplit {
		addrUint32 = addrUint32 << 8
		addrUint32 = addrUint32 | uint32(v)
	}

	return addrUint32, nil
}
