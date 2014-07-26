package httpd

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
)

// Logger prints the request method with url and then executes the next Handler.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}

// NotFound will send a context conform NotFound response.
func NotFound(w http.ResponseWriter, r *http.Request) {
	c := &context.Context{
		Request:  r,
		Response: w,
	}
	c.Error(nil, http.StatusNotFound)
}

// ErrorCode retrieves the correct http error code depending on the provided
// error type.
//
// NOTE: It may be an idea to add status code information to a specific Error
// type however this may couple to hard to a HTTP interface.
func ErrorCode(err error) int {
	c := http.StatusInternalServerError

	switch err {
	case mgo.ErrNotFound:
		c = http.StatusNotFound
	case common.ErrBadQueryId:
		c = http.StatusBadRequest
	}

	switch err.(type) {
	case validation.Error:
		c = http.StatusBadRequest
	}

	return c
}
