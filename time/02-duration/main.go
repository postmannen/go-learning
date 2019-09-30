package main

import (
	"fmt"
	"time"
)

func main() {
	duration := time.Duration(3) * time.Second

	for {
		<-time.After(duration)
		fmt.Println("monkey")
	}
}
