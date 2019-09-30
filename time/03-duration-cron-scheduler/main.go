package main

import (
	"fmt"
	"time"
)

func doSome1() {
	fmt.Println("first thing")
}

func doSome2() {
	fmt.Println("second thing")
}

func main() {
	duration1 := time.Duration(1) * time.Second
	duration2 := time.Duration(2) * time.Second

	go func(f func()) {
		for {
			<-time.After(duration1)
			f()

		}
	}(doSome1)

	go func(f func()) {
		for {
			<-time.After(duration2)
			f()

		}
	}(doSome2)

	<-time.After(time.Duration(10) * time.Second)
}
