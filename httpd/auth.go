package httpd

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/satisfeet/hoopoe/httpd/context"
)

type Auth struct {
	Username string
	Password string
	Handler  http.Handler
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := context.NewContext(w, r)
	h := c.Get("Authorization")

	if i := strings.IndexRune(h, ' '); i != -1 {
		b := []byte(a.Username + ":" + a.Password)

		if base64.StdEncoding.EncodeToString(b) == h[i+1:] {
			a.Handler.ServeHTTP(w, r)

			return
		}
	}

	c.Set("WWW-Authenticate", "Basic realm=hoopoe")
	c.Error(nil, http.StatusUnauthorized)
}
