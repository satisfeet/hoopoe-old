package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Handler struct {
	Auth    *handler.Auth
	Logger  *handler.Logger
	handler http.Handler
}

func NewHandler(s *store.Store) *Handler {
	m := http.NewServeMux()
	m.Handle("/customers", NewCustomerHandler(s))

	h := &Handler{
		Auth:    &handler.Auth{},
		Logger:  &handler.Logger{},
		handler: m,
	}
	h.handler = h.Logger.Handle(h.handler)
	h.handler = h.Auth.Handle(h.handler)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func ErrorCode(err error) int {
	c := http.StatusInternalServerError

	switch err {
	case mgo.ErrNotFound:
		c = http.StatusNotFound
	case mongo.ErrBadQueryParam:
		c = http.StatusBadRequest
	}

	switch err.(type) {
	case validation.Error:
		c = http.StatusBadRequest
	}

	return c
}
