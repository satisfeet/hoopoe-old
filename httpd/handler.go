package httpd

import (
	"log"
	"net/http"
)

const (
	JSON = "application/json"
)

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}

func Accept(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if a := r.Header.Get("Accept"); len(a) != 0 && a != JSON {
			Error(w, nil, http.StatusNotAcceptable)

			return
		}

		h.ServeHTTP(w, r)
	})
}

func ContentType(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if t := r.Header.Get("Content-Type"); len(t) != 0 && t != JSON {
			Error(w, nil, http.StatusUnsupportedMediaType)

			return
		}

		h.ServeHTTP(w, r)
	})
}
