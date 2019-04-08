package main

type aType struct {
	i int
}

func (a aType) aFn(i int) int {
	return i
}

type bType struct {
	i int
}

func (b *bType) bFn(i int) {
	b.i = i
}

func main() {

}
