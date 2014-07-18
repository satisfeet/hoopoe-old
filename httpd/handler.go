package httpd

import (
	"encoding/base64"
	"net/http"
	"strings"
)

var (
	HttpUsername = "bodokaiser"
	HttpPassword = "secret"
)

// Responds 401 if HTTP Basic authentication fails against
// HttpUsername and HttpPassword else it calls the wrapped
// http.Handler.
func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if v := strings.Split(r.Header.Get("Authorization"), " "); len(v) == 2 {
			if v, err := base64.StdEncoding.DecodeString(v[1]); err == nil {
				a := strings.Split(string(v), ":")

				if HttpUsername == a[0] && HttpPassword == a[1] {
					h.ServeHTTP(w, r)

					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", "Basic realm=hoopoe")

		Error(w, nil, http.StatusUnauthorized)
	})
}

// Responds 404.
func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Error(w, nil, http.StatusNotFound)
	})
}
