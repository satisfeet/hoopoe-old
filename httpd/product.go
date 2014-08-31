package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/store"
)

type ProductHandler struct {
	Store *store.ProductStore
}

func (h *ProductHandler) List(c *Context) {
	m := []store.Product{}

	q := store.NewProductQuery()

	if err := h.Store.Find(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *ProductHandler) Show(c *Context) {
	m := store.Product{}

	q := store.NewProductQuery()
	q.Where("id", c.Param("product"))

	if err := h.Store.FindOne(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}
