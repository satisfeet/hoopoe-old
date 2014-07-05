package httpd

import (
	"errors"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, e error, c int) {
	if e == nil {
		e = errors.New(http.StatusText(c))
	}

	log.Printf("Error: %s", e.Error())

	http.Error(w, `{"error":"`+e.Error()+`"}`, c)
}
