package httpd

import (
	"io"
	"net/http"

	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/common"
)

type ProductHandler struct {
	store *store.ProductStore

	Image *ProductImageHandler
}

func NewProductHandler(s *common.Session) (*ProductHandler, error) {
	store, err := store.NewProductStore(s)

	return &ProductHandler{
		store: store,
		Image: &ProductImageHandler{
			store: store,
		},
	}, err
}

func (h *ProductHandler) List(c *Context) {
	m := []store.Product{}

	q := store.NewProductQuery()

	if err := h.store.FindAll(q, &m); err != nil {
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

type ProductImageHandler struct {
	store *store.ProductStore
}

func (h *ProductImageHandler) Show(c *Context) {
	m := store.Product{}

	q := store.NewProductQuery()
	q.Id(c.Param("product"))

	if err := h.store.FindOne(q, &m); err != nil {
		c.Error(err)

		return
	}

	f, err := h.store.Image.Open(m.Id)

	if err != nil {
		c.Error(err)

		return
	}

	defer f.Close()

	if _, err := io.Copy(c.Response, f); err != nil {
		c.Error(err)
	}
}

func (h *ProductImageHandler) Update(c *Context) {
	m := store.Product{}

	q := store.NewProductQuery()
	q.Id(c.Param("product"))

	if err := h.store.FindOne(q, &m); err != nil {
		c.Error(err)

		return
	}

	h.store.Image.Remove(m.Id)

	f, err := h.store.Image.New(m.Id)

	if err != nil {
		c.Error(err)

		return
	}

	defer f.Close()

	if _, err := io.Copy(f, c.Request.Body); err != nil {
		c.Error(err)

		return
	}

	if err := h.store.Update(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}
