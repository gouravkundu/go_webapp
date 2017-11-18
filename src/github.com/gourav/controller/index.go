package controller

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gourav/util"
)

var tpl *template.Template

type Person struct {
	First_name, Last_name util.Name
}

func (p Person) Full_name() string {
	name := fmt.Sprintf("%s %s", p.First_name.ToCapitalize(), p.Last_name.ToCapitalize())
	return name
}

func (p Person) IsValid() bool {
	if p == (Person{}) {
		fmt.Println("Person struct not present")
		return false
	}
	return true
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}

func PassingFormParsing(w http.ResponseWriter, r *http.Request) {
	var person Person
	if r.Method == http.MethodPost {
		person = Person{
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
