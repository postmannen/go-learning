package main

import (
	"fmt"
	"time"
)

func main() {
	highPri := make(chan int)
	lowPri := make(chan int)

	go func() {

		for i := 0; i <= 5; i++ {
			highPri <- i
			time.Sleep(time.Millisecond * 250)
		}
		close(highPri)
		fmt.Println("closing highPri")
	}()

	go func() {

		for i := 0; i <= 5; i++ {
			lowPri <- i
			time.Sleep(time.Millisecond * 501)
		}
		close(lowPri)
		fmt.Println("closing lowPri")
	}()

	for {
		select {
		case i, ok := <-highPri:
			fmt.Println("High Pri", i)
			if !ok {
				highPri = nil
			}
			continue
		default:
			select {
			case i, ok := <-highPri:
				fmt.Println("High Pri", i)
				if !ok {
					highPri = nil
				}
			case i, ok := <-lowPri:
				fmt.Println("Low Pri", i)
				if !ok {
					lowPri = nil
				}
			}
		}

		if lowPri == nil && highPri == nil {
			break
		}
	}

}

