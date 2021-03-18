package main

import (
	"bytes"
	"os"
	"text/template"
)

func main() {
	t := template.New("")

	render := func(name string) (string, error) {
		var buf bytes.Buffer
		err := t.ExecuteTemplate(&buf, name, nil)
		return buf.String(), err
	}

	funcs := template.FuncMap{
		"render": render,
	}

	str := `
		Define T1             {{define "T1"}}Hello{{end}} 
		Capture a variable    {{$v := render "T1"}} 
		Call len or something {{len $v}}
	`

	tt, err := t.Funcs(funcs).Parse(str)
	if err != nil {
		panic(err)
	}

	if err := tt.Execute(os.Stdout, nil); err != nil {
		panic(err)
	}
}
