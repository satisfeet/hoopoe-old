package httpd

import (
	"database/sql"
	"net/http"

	"github.com/satisfeet/hoopoe/store"
)

type ProductHandler struct {
	store *store.ProductStore
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		store: store.NewProductStore(db),
	}
}

func (h *ProductHandler) List(c *Context) {
	m := []store.Product{}

	if err := h.store.Find(&m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *ProductHandler) Show(c *Context) {
	m := store.Product{}

	if err := h.store.FindId(c.Param("product"), &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}
