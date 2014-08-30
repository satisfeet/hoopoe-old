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

	if err := h.Store.Find(&m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}
