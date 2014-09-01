package httpd

import (
	"database/sql"
	"net/http"

	"github.com/satisfeet/hoopoe/store"
)

type CustomerHandler struct {
	store *store.CustomerStore
}

func NewCustomerHandler(db *sql.DB) *CustomerHandler {
	return &CustomerHandler{
		store: store.NewCustomerStore(db),
	}
}

func (h *CustomerHandler) List(c *Context) {
	m := []store.Customer{}

	if err := h.store.Search(c.Query("search"), &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Show(c *Context) {
	m := store.Customer{}

	if err := h.store.FindId(c.Param("customer"), &m); err != nil {
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

	if err := h.store.Insert(&m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Update(c *Context) {
	m := store.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err)

		return
	}

	if err := h.store.UpdateId(c.Param("customer"), &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}

func (h *CustomerHandler) Destroy(c *Context) {
	if err := h.store.RemoveId(c.Param("customer")); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}
