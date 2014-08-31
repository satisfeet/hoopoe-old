package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/store"
)

type CustomerHandler struct {
	Store *store.CustomerStore
}

func (h *CustomerHandler) List(c *Context) {
	m := []store.Customer{}

	q := store.NewCustomerQuery()

	if err := h.Store.Find(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Show(c *Context) {
	m := store.Customer{}

	q := store.NewCustomerQuery()
	q.Where("id", c.Param("customer"))

	if err := h.Store.FindOne(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}
