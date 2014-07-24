package httpd

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/satisfeet/hoopoe/httpd/context"
)

// Auth is a http.Handler which wraps a http.Handler for requests which do not
// match the required authentication.
type Auth struct {
	Username string
	Password string
	Handler  http.Handler
}

// Implementation of the http.Handler interface. Validates the HTTP Request
// against HTTP Basic.
//
// NOTE: To be only used over HTTPS to not expose credentials!
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
