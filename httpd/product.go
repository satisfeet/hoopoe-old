package httpd

import (
	"database/sql"
	"net/http"

	"github.com/satisfeet/hoopoe/model"
)

type ProductHandler struct {
	store *model.ProductStore
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		store: model.NewProductStore(db),
	}
}

func (h *ProductHandler) List(c *Context) {
	m := []model.Product{}

	if err := h.store.Find(&m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *ProductHandler) Show(c *Context) {
	m := model.Product{}

	if err := h.store.FindId(c.Param("product"), &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}
