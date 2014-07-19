package httpd

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/satisfeet/hoopoe/httpd/context"
)

var (
	Username = "bodokaiser"
	Password = "secret"
)

func Auth(h http.Handler) http.Handler {
	b := base64.StdEncoding.EncodeToString([]byte(Username + ":" + Password))

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
