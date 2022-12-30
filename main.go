package main

import (
	"html/template"

	"photo.dev/app"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("assets/templates/*"))
}

func main() {
	Application := app.App{}
	Application.Init(tpl)
	Application.Run()
}