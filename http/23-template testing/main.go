package main

import (
	"html/template"
	"log"
	"os"
)

type user struct {
	Name string
}

func main() {
	tf, err := template.ParseFiles("page.html")
	if err != nil {
		log.Println("Error parsing template", err)
	}

	myUser := user{
		Name: "Bjørn Tore",
	}

	tf.Execute(os.Stdout, myUser)

}
