package httpd

import (
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Customer struct {
	store *store.Customer
}

func NewCustomer(s *mongo.Store) *Customer {
	return &Customer{
		store: store.NewCustomer(s),
	}
}

func (h *Customer) List(c *context.Context) {
	m := []model.Customer{}

	if err := h.store.Search(c.Query("search"), &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customer) Show(c *context.Context) {
	m := model.Customer{}

	if err := h.store.FindId(c.Param("customer"), &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customer) Create(c *context.Context) {
	m := model.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Insert(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customer) Update(c *context.Context) {
	m := model.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Update(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customer) Destroy(c *context.Context) {
	if err := h.store.RemoveId(c.Param("customer")); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}
