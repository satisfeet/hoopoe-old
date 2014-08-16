package httpd

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-validation"

	"github.com/satisfeet/hoopoe/store/mongo"
)

type Context struct {
	*context.Context
}

func (c *Context) Error(err error) {
	s := http.StatusInternalServerError

	switch err {
	case mgo.ErrNotFound:
		s = http.StatusNotFound
	case mongo.ErrBadId:
		s = http.StatusBadRequest
	}

	switch err.(type) {
	case *json.UnmarshalTypeError, validation.Error:
		s = http.StatusBadRequest
	}

	c.Context.Error(err, s)
}

type HandlerFunc func(*Context)

func (f HandlerFunc) ServeHTTP(c *context.Context) {
	f(&Context{Context: c})
}
