package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gourav/controller"
)

func main() {
	http.HandleFunc("/", moveToIndex)
	http.Handle("/fabicon.ico", http.NotFoundHandler())
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/file", controller.FileFormPosting)
	http.HandleFunc("/valuepassing", controller.PassingUrlValue)
	http.HandleFunc("/passingformvalue", controller.PassingFormParsing)
	http.HandleFunc("/set", setCookieCustom)
	http.HandleFunc("/get", getCookieCustom)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setCookieCustom(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "123456",
	})
	io.WriteString(w, "Cookie has been set on the client")
}

func getCookieCustom(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	value := fmt.Sprintf("Value of the stored %s at the client end is: %s", c.Name, c.Value)
	io.WriteString(w, value)
}

func moveToIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Redirecting to /index page")
	//	w.Header().Set("Location", "/index")
	//	w.WriteHeader(http.StatusMovedPermanently)
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
