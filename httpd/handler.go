package httpd

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type handler struct{}

func (h *handler) error(c *context.Context, err error) {
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

	c.Error(err, s)
}
