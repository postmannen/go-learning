package views

import (
	"html/template"
)

//NewView used to create new view with templates attached
func NewView(layout string, files ...string) *View {
	files = append(files,
		"views/layouts/bootstrap.html",
		"views/layouts/navbar.html",
		"views/layouts/footer.html",
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

//View the type
type View struct {
	Template *template.Template
	Layout   string
}
