package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second * 1)
	fmt.Println(time.Until(start))
}
