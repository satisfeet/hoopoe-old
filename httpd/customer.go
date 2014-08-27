package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/common"
)

type CustomerHandler struct {
	store *store.CustomerStore
}

func NewCustomerHandler(s *common.Session) (*CustomerHandler, error) {
	store, err := store.NewCustomerStore(s)

	return &CustomerHandler{
		store: store,
	}, err
}

func (h *CustomerHandler) List(c *Context) {
	m := []store.Customer{}

	q := store.NewCustomerQuery()
	q.Search(c.Query("search"))

	if err := h.store.FindAll(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *CustomerHandler) Show(c *Context) {
	m := store.Customer{}

	q := store.NewCustomerQuery()
	q.Id(c.Param("customer"))

	if err := h.store.FindOne(q, &m); err != nil {
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

	q := store.NewCustomerQuery()
	q.Id(c.Param("id"))

	if err := c.Parse(&m); err != nil {
		c.Error(err)

		return
	}

	if err := h.store.Update(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}

func (h *CustomerHandler) Destroy(c *Context) {
	q := store.NewCustomerQuery()
	q.Id(c.Param("id"))

	if err := h.store.Remove(q); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}
