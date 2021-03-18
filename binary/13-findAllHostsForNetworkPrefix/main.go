package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

// Will take a prefix and mask like 192.160.0.0/24 as input, and
// return all the allowed host adresses for that prefix.
func cidrHosts(prefixAndMask string) ([]string, error) {
	prefixAndMaskSplit := strings.Split(prefixAndMask, "/")
	prefixString := prefixAndMaskSplit[0]

	// Create a byte slice of the dot'ed ip address.
	prefixByteSlice, err := convertDotStringToByteSlice(prefixString)
	if err != nil {
		return nil, err
	}

	// Create an uint32 value of the ip address.
	var prefixUint32 uint32
	for _, v := range prefixByteSlice {
		prefixUint32 = prefixUint32 << 8
		prefixUint32 = prefixUint32 | uint32(v)
	}

	// Convers the mask bits to int
	maskBitsNRString := prefixAndMaskSplit[1]
	maskBitsNRInt, err := strconv.Atoi(maskBitsNRString)
	if err != nil {
		return nil, fmt.Errorf("error: failed to convert maskString to int: %v", err)
	}

	hostBitsNRInt := 32 - maskBitsNRInt
	hostsNRAllowed := int(math.Pow(2, float64(hostBitsNRInt))) - 1
	hostIPs := []string{}

	// Loop and create the host portion of the address for the
	// number of allowed hosts based on the host part of the mask.
	for i := 0; i <= hostsNRAllowed; i++ {
		ha := prefixUint32 | uint32(i)
		hIP := convertUint32ToDotedString(ha)
		hostIPs = append(hostIPs, hIP)
	}

	return hostIPs, nil
}

// Will convert string x.x.x.x of ip address into
// a byte slice of 4 elements.
func convertDotStringToByteSlice(s string) ([]byte, error) {
	nSplit := make([]byte, 4)

	sSplit := strings.Split(s, ".")
	for _, v := range sSplit {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("error: strconv.Atoi failed: %v", err)
		}

		nSplit = append(nSplit, byte(n))
	}

	return nSplit, nil
}

// Convert the uint32 representation of an ip address into
// a x.x.x.x string representation.
func convertUint32ToDotedString(u uint32) string {
	bs := make([]byte, 4)
	lsb := uint32(0x000000ff)

	for i := 3; i >= 0; i-- {
		b := byte(u & lsb)
		u = u >> 8
		bs[i] = b
	}

	ipString := fmt.Sprintf("%v.%v.%v.%v", bs[0], bs[1], bs[2], bs[3])

	return ipString
}

func main() {
	prefixAndMask := "127.0.255.4/30"
	allowedHosts, err := cidrHosts(prefixAndMask)
	if err != nil {
		log.Printf("%v\n", err)
	}

	fmt.Printf("Main: Printing out all the hosts: %v\n", allowedHosts)
}
