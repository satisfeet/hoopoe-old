package httpd

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/common"
)

func Handler(u, p string) http.Handler {
	c := (&Customers{
		Store: &store.Customers{
			Mongo: store.DefaultMongo,
		},
	}).Handler()

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/customers"):
			c.ServeHTTP(w, r)
		default:
			handler.NotFound(w, r)
		}
	})

	a := &handler.Auth{
		Username: u,
		Password: p,
		Handler:  h,
	}

	return &handler.Logger{a}
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
