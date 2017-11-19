package controller

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	User_id    string
	First_name string
	Last_name  string
}

func (u user) IsValid() bool {
	if u != (user{}) {
		return true
	} else {
		return false
	}
}

//var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

// var dbSessions = make(map[string]string)

//func init() {
//	tpl = template.Must(template.ParseGlob("./templates/*"))
//}

func Foo(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	//Check if the user already exists.
	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}

	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbUsers[un] = u
		dbSessions[cookie.Value] = un
	}

	err = tpl.ExecuteTemplate(w, "session.go.html", u)
	handleError(w, err)
}

func Bar(w http.ResponseWriter, r *http.Request) {
	var u user
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("Cookie is not set yet!")
	}
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	err = tpl.ExecuteTemplate(w, "session_details.go.html", u)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
