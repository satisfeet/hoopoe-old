package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/common"
)

type ProductHandler struct {
	store *store.ProductStore
}

func NewProductHandler(s *common.Session) (*ProductHandler, error) {
	store, err := store.NewProductStore(s)

	return &ProductHandler{
		store: store,
	}, err
}

func (h *ProductHandler) List(c *Context) {
	m := []store.Product{}

	if err := h.store.FindAll(nil, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *ProductHandler) Show(c *Context) {
	m := store.Product{}

	q := store.NewProductQuery()
	q.Id(c.Param("product"))

	if err := h.store.FindOne(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *ProductHandler) Create(c *Context) {
	m := store.Product{}

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

func (h *ProductHandler) Update(c *Context) {
	m := store.Product{}

	q := store.NewProductQuery()
	q.Id(c.Param("product"))

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

func (h *ProductHandler) Destroy(c *Context) {
	q := store.NewProductQuery()
	q.Id(c.Param("product"))

	if err := h.store.Remove(q); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}
