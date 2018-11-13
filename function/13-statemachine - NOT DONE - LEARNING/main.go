//The idea here is to test out having a function returning
// another function to be executed next.
// The reason we might want to do something like this can
// be that we are iterating a slice, and if we want to do
// something special depending on the value found, we return
// the next function to be called. That function again can
// call another one...and so on.
package main

import "fmt"

type aFunction func() aFunction

type data struct {
	counter int
	length  int
	numbers []int
	aFunction
}

func (d *data) less() aFunction {
	fmt.Println("less : ", d.counter, " is less than 5")
	// less is the last function in the chain, so we call the
	// initial d.checkSize function again to get the next
	// number, and start all over.
	return d.checkSize
}

func (d *data) more() aFunction {
	fmt.Println("more : ", d.counter, " is more than 5")
	// less is the last function in the chain, so we call the
	// initial d.checkSize function again to get the next
	// number, and start all over.
	return d.checkSize
}

func (d *data) checkSize() aFunction {
	d.counter++
	// If we are at the end of the slice, return nil back to main,
	// and break the for loop there. Since nil contains no function
	// to call no more functions will be called.
	if d.counter == d.length {
		return nil
	}

	// If the number is below 5, return the d.less function to be
	// executed in main for loop.
	if d.numbers[d.counter] <= 5 {
		fmt.Println("checkSize : ", d.counter, " is less than 5")
		return d.less
	}

	// If the number is above 5, return the d.more function to be
	// executed in main for loop.
	if d.numbers[d.counter] > 5 {
		fmt.Println("checkSize : ", d.counter, " is larger than 5")
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

	// Here we are starting the first function who will start it all,
	// and put the returned function into f.
	f := d.checkSize()
	for {
		fmt.Println("------------------------------------------------------------------")
		fmt.Printf("Main: Calling %#v, type = %T\n", f, f)
		// Here we are constantly looping and executing the returned functions.
		// Again we put the returned function value into f to then be executed
		// on the next round of the for loop.
		// The d.checkSize function will check the if we have reached the end of
		// the slice, and return nil instead of another function to call so we
		// can exit the for loop here by checking for nil, and then break the loop.
		f = f()
		if f == nil {
			break
		}

	}

}
