package controller

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

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

func FileFormPosting(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		f, h, e := r.FormFile("file_struct")
		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(r)
		fmt.Println(h)
		fmt.Println(f)
		fmt.Println(r.FormValue("fname"))
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
		file, err := os.Create(filepath.Join("./container/", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = file.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
	err := tpl.ExecuteTemplate(w, "file_form_parsing.go.html", s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PassingFormParsing(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	var err error
	if r.Method == http.MethodPost {
		person = models.Person{
			First_name: util.Name(r.FormValue("fname")),
			Last_name:  util.Name(r.FormValue("lname")),
		}
		err = tpl.ExecuteTemplate(w, "form_parsing.go.html", person)
	} else {
		err = tpl.ExecuteTemplate(w, "form_parsing.go.html", nil)
	}
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
