package main

import "fmt"

func main() {
	outer := struct {
		inner struct {
			n string
		}
	}{
		inner: struct{ n string }{n: "hest"},
	}

	fmt.Println(outer.inner.n)

}
