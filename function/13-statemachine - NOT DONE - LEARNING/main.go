package main

import "fmt"

type aFunction func() aFunction

type data struct {
	counter int
	length  int
	numbers []int
}

func (d *data) someFunction() aFunction {
	d.counter++
	if d.counter == d.length {
		return nil
	}

	if d.numbers[d.counter] < 5 {
		fmt.Println(d.counter, " is less than 5")
		return d.someFunction
	}
	if d.numbers[d.counter] > 5 {
		fmt.Println(d.counter, " is larger than 5")
		return d.someFunction
	}

	return nil
}

func main() {
	d := &data{
		counter: 0,
		numbers: []int{0, 1, 2, 4, 5, 6, 7, 8, 9},
	}
	d.length = len(d.numbers)

	d.someFunction()

}
