package main

type fakeDB struct {
	column1 string
}

func (fd *fakeDB) read(p []byte) (n int, err error) {
	fd.column1 = string(p)
	return
}

func main() {

}
