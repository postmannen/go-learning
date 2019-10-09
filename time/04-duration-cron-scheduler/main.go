package main

import (
	"fmt"
	"runtime"
	"time"
)

func doSome1() {
	fmt.Println("first thing")
}

func doSome2() {
	fmt.Println("second thing")
}

type toSchedule struct {
	fønk     func()
	duration time.Duration
}

func main() {
	runtime.GOMAXPROCS(1)
	schedulePlan := []toSchedule{
		toSchedule{fønk: doSome1, duration: time.Duration(1) * time.Second},
		toSchedule{fønk: doSome2, duration: time.Duration(2) * time.Second},
	}

	for _, s := range schedulePlan {
		go func(s toSchedule) {
			for {
				<-time.After(s.duration)
				s.fønk()

			}
		}(s)
	}

	<-time.After(time.Duration(10) * time.Second)
}
