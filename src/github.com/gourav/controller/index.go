package controller

import (
	"html/template"
	"io"
	"net/http"

	"github.com/gourav/models"
	"github.com/gourav/util"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func PassingFormParsing(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	if r.Method == http.MethodPost {
		person = models.Person{
			First_name: util.Name(r.FormValue("fname")),
			Last_name:  util.Name(r.FormValue("lname")),
		}
	}
	err := tpl.ExecuteTemplate(w, "form_parsing.go.html", person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PassingUrlValue(w http.ResponseWriter, r *http.Request) {
	search_data := r.FormValue("q")
	err := tpl.ExecuteTemplate(w, "passing_value.go.html", search_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
