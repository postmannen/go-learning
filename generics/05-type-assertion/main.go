package main

import (
	"fmt"
	"strconv"
	"time"
)

type Scalar interface {
	int | float64 | string | time.Time
}

func scalarFromString[S Scalar](val string) (r S, err error) {
	switch rp := any(&r).(type) {
	case *int:
		*rp, err = strconv.Atoi(val)
	case *string:
		*rp = val
	case *float64:
		*rp, err = strconv.ParseFloat(val, 64)
	case *time.Time:
		*rp, err = time.Parse(time.RFC3339, val)
	default:
		panic("unreachable")
	}
	return
}

func main() {
	fmt.Println(scalarFromString[float64]("8.2"))
	fmt.Println(scalarFromString[time.Time]("bad time"))
}
