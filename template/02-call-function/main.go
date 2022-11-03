package main

import (
	"os"
	"text/template"
)

func main() {

	t := `{{ horse 32 }}`

	echo := func(in any) any {
		return in
	}

	funcs := template.FuncMap{
		"horse": echo,
	}

	tmpl, err := template.New("test").Funcs(funcs).Parse(t)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
