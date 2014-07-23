package httpd

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/httpd/context"
	"github.com/satisfeet/hoopoe/model/validation"
	"github.com/satisfeet/hoopoe/store"
)

// Logger prints the request method with url and then executes
// the next http.Handler.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}

// NotFound will send a context conform NotFound response.
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
