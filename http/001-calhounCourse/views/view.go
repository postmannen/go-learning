package views

import (
	"fmt"
	"html/template"
	"path/filepath"
)

//LayoutDir the directory where all the layouts are, defaults to "views/layouts/""
var LayoutDir = "views/layouts/"

//TemplateExt defaults to ".html"
var TemplateExt = ".html"

//NewView used to create new view with templates attached
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
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

//layoutFiles will get a slice containing the name of all the *.html files in the views/layout directory
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		fmt.Println("Error: filepath.Glob failed :", err)
	}
	return files
}
