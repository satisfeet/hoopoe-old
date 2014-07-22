package httpd

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/satisfeet/hoopoe/httpd/context"
)

var (
	Basic = ""
)

func Auth(h http.Handler) http.Handler {
	b := base64.StdEncoding.EncodeToString([]byte(Basic))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := context.NewContext(w, r)
		a := c.Get("Authorization")

		if i := strings.IndexRune(a, ' '); i != -1 {
			if b == a[i+1:] {
				h.ServeHTTP(w, r)
				return
			}
		}

		c.Set("WWW-Authenticate", "Basic realm=hoopoe")
		c.Error(nil, http.StatusUnauthorized)
	})
}

func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.NewContext(w, r).Error(nil, http.StatusNotFound)
	})
}
