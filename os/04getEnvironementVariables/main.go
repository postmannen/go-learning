package main

import (
	"fmt"
	"os"
)

func main() {
	//export MY_ENV="APEKATT"
	env := os.Getenv("MY_ENV")
	fmt.Println(env)

}
