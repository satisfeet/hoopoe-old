package httpd

import (
	"log"
	"net/http"
	"strings"
)

const (
	TYPE_ALL  = "*/*"
	TYPE_JSON = "application/json"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}

func Accept(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a := r.Header.Get("Accept"); len(a) != 0 && !(contains(a, TYPE_JSON) || contains(a, TYPE_ALL)) {
			Error(w, nil, http.StatusNotAcceptable)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func ContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if t := r.Header.Get("Content-Type"); len(t) != 0 && t != TYPE_JSON {
			Error(w, nil, http.StatusUnsupportedMediaType)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func contains(h string, p string) bool {
	return strings.Contains(h, p)
}
