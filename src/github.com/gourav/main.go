package main

import (
	"log"
	"net/http"

	"github.com/gourav/controller"
)

func main() {
	http.HandleFunc("/index", controller.Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
