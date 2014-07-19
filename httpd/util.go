package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/store"
)

// ErrorCode retrieves the correct http error code
// depending on the provided error type.
func ErrorCode(err error) int {
	c := http.StatusInternalServerError

	switch err {
	case mgo.ErrNotFound:
		c = http.StatusNotFound
	case store.ErrInvalidQuery:
		c = http.StatusBadRequest
	}

	return c
}
