package httpd

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
)

var DefaultCustomers = NewCustomers()

func Handler(u, p string) http.Handler {
	h := http.HandlerFunc(mux)

	a := &handler.Auth{
		Username: u,
		Password: p,
		Handler:  h,
	}

	return &handler.Logger{a}
}

func mux(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/customers"):
		DefaultCustomers.ServeHTTP(w, r)
	default:
		handler.NotFound(w, r)
	}
}

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
