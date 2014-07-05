package httpd

import (
	"log"
	"net/http"
)

func Accept(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			Error(w, nil, http.StatusNotAcceptable)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}
