package httpd

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/httpd/context"
	"github.com/satisfeet/hoopoe/model/validation"
	"github.com/satisfeet/hoopoe/store"
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

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	context.NewContext(w, r).Error(nil, http.StatusNotFound)
}

// ErrorCode retrieves the correct http error code
// depending on the provided error type.
func ErrorCode(err error) int {
	c := http.StatusInternalServerError

	switch err {
	case mgo.ErrNotFound, store.ErrInvalidQuery:
		c = http.StatusNotFound
	}

	switch err.(type) {
	case validation.Error:
		c = http.StatusBadRequest
	}

	return c
}
