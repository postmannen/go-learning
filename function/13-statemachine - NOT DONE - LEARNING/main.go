package main

import "fmt"

type aFunction func() aFunction

type data struct {
	counter int
	length  int
	numbers []int
}

func (d *data) less() aFunction {
	fmt.Println("less : ", d.counter, " is less than 5")
	return d.someFunction
}

func (d *data) more() aFunction {
	fmt.Println("more : ", d.counter, " is more than 5")
	return d.someFunction
}

func (d *data) someFunction() aFunction {
	d.counter++
	if d.counter == d.length {
		return nil
	}

	if d.numbers[d.counter] <= 5 {
		fmt.Println("someFunction : ", d.counter, " is less than 5")
		return d.less
	}
	if d.numbers[d.counter] > 5 {
		fmt.Println("someFunction : ", d.counter, " is larger than 5")
		return d.more
	}

	return nil
}

func main() {
	d := &data{
		counter: 0,
		numbers: []int{0, 1, 2, 4, 5, 6, 7, 8, 9, 10, 11},
	}
	d.length = len(d.numbers)
	for {
		f := d.someFunction()
		fmt.Println("DEBUG : in for loop d.counter = ", d.counter)
		if f == nil {
			break
		}

	}

}
