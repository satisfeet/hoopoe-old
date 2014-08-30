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

	if err := h.Store.Find(&m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}
