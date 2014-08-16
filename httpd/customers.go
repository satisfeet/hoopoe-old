package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-context"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Customer struct {
	*handler
	store *store.Customer
}

func NewCustomer(s *mgo.Session) *Customer {
	return &Customer{
		store: store.NewCustomer(s),
	}
}

func (h *Customer) List(c *context.Context) {
	m := []model.Customer{}

	if err := h.store.Search(c.Query("search"), &m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customer) Show(c *context.Context) {
	m := model.Customer{}
	m.Id = store.IdFromString(c.Param("customer"))

	if err := h.store.FindOne(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customer) Create(c *context.Context) {
	m := model.Customer{}

	if err := c.Parse(&m); err != nil {
		h.error(c, err)

		return
	}

	if err := h.store.Insert(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customer) Update(c *context.Context) {
	m := model.Customer{}

	if err := c.Parse(&m); err != nil {
		h.error(c, err)

		return
	}

	if err := h.store.Update(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customer) Destroy(c *context.Context) {
	m := model.Customer{}
	m.Id = store.IdFromString(c.Param("customer"))

	if err := h.store.Remove(m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}
