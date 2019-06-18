package somepkg

import (
	"encoding/json"
	"fmt"
	"log"
)

//MarshallSome will marshall some
func MarshallSome() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "Bob",
		Age:  100,
	}

	b, err := json.Marshal(a)
	if err != nil {
		log.Println("error: marshalling failed: ", err)
	}
	fmt.Println(string(b))
}
