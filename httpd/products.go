package httpd

import (
	"io"
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Product struct {
	store *store.Product
}

func NewProduct(s *mongo.Store) *Product {
	return &Product{
		store: store.NewProduct(s),
	}
}

func (h *Product) List(c *context.Context) {
	m := []model.Product{}

	if err := h.store.Find(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Product) Show(c *context.Context) {
	m := model.Product{}

	if err := h.store.FindId(c.Param("product"), &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Product) Create(c *context.Context) {
	m := model.Product{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Insert(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Product) Update(c *context.Context) {
	m := model.Product{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Update(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Product) Destroy(c *context.Context) {
	if err := h.store.RemoveId(c.Param("product")); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Product) ShowImage(c *context.Context) {
	f, err := h.store.OpenImage(c.Param("product"), c.Param("image"))
	if err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		io.Copy(c.Response, f)
	}
}

func (h *Product) CreateImage(c *context.Context) {
	f, err := h.store.CreateImage(c.Param("product"))
	if err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	if _, err := io.Copy(f, c.Request.Body); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Product) DestroyImage(c *context.Context) {
	if err := h.store.RemoveImage(c.Param("product"), c.Param("image")); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}