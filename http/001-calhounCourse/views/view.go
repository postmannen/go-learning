package views

import (
	"html/template"
)

//NewView used to create new view with templates attached
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.html")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

//View the type
type View struct {
	Template *template.Template
}
