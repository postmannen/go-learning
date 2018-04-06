package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//Write method wants a []byte, so we typecast a string into a []byte
		w.Write([]byte(`
		<html> 
        	<head> 
          		<title>Chat</title> 
        	</head> 
        	<body> 
          		Let's chat! 
        	</body> 
      	</html>
		`))
	})

	//we can bundle the output for checking directly in calling a function.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error : ListenAndServe failed : ", err)
	}
}
