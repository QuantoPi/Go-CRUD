package routes

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))
	data := ""
	tmpl.Execute(w, data)
}
func CostumersHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/costumers.gohtml"))
	data := ""
	tmpl.Execute(w, data)
}
func DatabaseHandler(w http.ResponseWriter, r *http.Request) {
	tmpl = template.Must(template.ParseFiles("templates/db-control.gohtml"))
	data := ""
	tmpl.Execute(w, data)
}
