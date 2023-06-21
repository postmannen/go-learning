package main

import (
	"html/template"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.Name}}\n")
	if err != nil {
		panic(err)
	}

	fh, err := os.Open("data.yaml")
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	in, err := io.ReadAll(fh)
	if err != nil {
		log.Fatalf("%v\n", err)
	}

	m := make(map[string]any)
	yaml.Unmarshal(in, &m)

	t1.ExecuteTemplate(os.Stdout, "t1", m)

}
