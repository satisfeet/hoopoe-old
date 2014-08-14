package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Product struct {
	*handler
	store *store.Product
}

func NewProduct(s *mgo.Session) *Product {
	return &Product{
		store: store.NewProduct(s),
	}
}

func (h *Product) List(c *context.Context) {
	m := []model.Product{}

	if err := h.store.Find(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Product) Show(c *context.Context) {
	m := model.Product{}
	m.Id = store.IdFromString(c.Param("product"))

	if err := h.store.FindOne(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Product) Create(c *context.Context) {
	m := model.Product{}

	if err := c.Parse(&m); err != nil {
		h.error(c, err)

		return
	}

	if err := h.store.Insert(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Product) Update(c *context.Context) {
	m := model.Product{}

	if err := c.Parse(&m); err != nil {
		h.error(c, err)

		return
	}

	if err := h.store.Update(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Product) Destroy(c *context.Context) {
	m := model.Product{}
	m.Id = store.IdFromString(c.Param("product"))

	if err := h.store.Remove(m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Product) ShowImage(c *context.Context) {
	id := store.IdFromString(c.Param("image"))

	m := model.Product{}
	m.Id = store.IdFromString(c.Param("product"))

	if err := h.store.ReadImage(&m, id, c.Response); err != nil {
		h.error(c, err)
	}
}

func (h *Product) CreateImage(c *context.Context) {
	id := bson.NewObjectId()

	m := model.Product{}
	m.Id = store.IdFromString(c.Param("product"))

	if err := h.store.WriteImage(&m, id, c.Request.Body); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Product) DestroyImage(c *context.Context) {
	id := store.IdFromString(c.Param("image"))

	m := model.Product{}
	m.Id = store.IdFromString(c.Param("product"))

	if err := h.store.RemoveImage(&m, id); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}
