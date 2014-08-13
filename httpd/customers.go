package httpd

import (
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Customer struct {
	*handler

	store *store.Customer
}

func NewCustomer(m *mongo.Store) *Customer {
	s := store.NewCustomer(m)

	return &Customer{
		handler: &handler{s},
		store:   s,
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

	if err := h.store.FindId(c.Param("customer"), &m); err != nil {
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
	m.Id = mongo.IdFromString(c.Param("customer"))

	if err := h.store.Remove(m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}
