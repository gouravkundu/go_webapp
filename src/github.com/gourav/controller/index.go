package controller

import (
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world")
}
