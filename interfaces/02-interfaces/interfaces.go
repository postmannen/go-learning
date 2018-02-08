package main

import "fmt"

type walker interface {
  walk(miles int)
}

type camel struct{
  Name string
}

func (c camel) walk(miles int) {
  fmt.Println(c.Name, "is walking ", miles)
}

func longWalk(w walker) {
  w.walk(500)
  w.walk(500)
}

func main() {
  c:=camel{"Bill"}
  longWalk(c)
}
