package main

import (
	"log"
	"net/http"

	"github.com/gourav/controller"
)

func main() {
	http.Handle("/fabicon.ico", http.NotFoundHandler())
	http.HandleFunc("/index", controller.Index)
	http.HandleFunc("/file", controller.FileFormPosting)
	http.HandleFunc("/valuepassing", controller.PassingUrlValue)
	http.HandleFunc("/passingformvalue", controller.PassingFormParsing)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
