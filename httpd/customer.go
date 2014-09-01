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

	if err := h.Store.Search(c.Query("search"), &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Show(c *Context) {
	m := store.Customer{}

	if err := h.Store.FindId(c.Param("customer"), &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Create(c *Context) {
	m := store.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err)

		return
	}

	if err := h.Store.Insert(&m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Destroy(c *Context) {
	if err := h.Store.RemoveId(c.Param("customer")); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}
