package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type group struct {
	Name string `validate:"startswith=g_"`
}

func main() {
	val := validator.New()

	{
		s := "apekatt"

		err := val.Var(s, "startswith=g_")
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}

	{
		g := group{
			Name: "ag_1",
		}

		err := val.Struct(g)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
	}
}
