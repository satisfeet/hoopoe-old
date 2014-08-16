package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-context"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Order struct {
	*handler
	store *store.Order
}

func NewOrder(s *mgo.Session) *Order {
	return &Order{
		store: store.NewOrder(s),
	}
}

func (h *Order) List(c *context.Context) {
	m := []model.Order{}

	if err := h.store.Find(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Order) Show(c *context.Context) {
	m := model.Order{}
	m.Id = store.IdFromString(c.Param("order"))

	if err := h.store.FindOne(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Order) Create(c *context.Context) {
	m := model.Order{}

	if err := c.Parse(&m); err != nil {
		h.error(c, err)
	}

	if err := h.store.Insert(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Order) Destroy(c *context.Context) {
	m := model.Order{}
	m.Id = store.IdFromString(c.Param("order"))

	if err := h.store.Remove(&m); err != nil {
		h.error(c, err)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Order) ShowInvoice(c *context.Context) {
	m := model.Order{}
	m.Id = store.IdFromString(c.Param("order"))

	if err := h.store.ReadInvoice(&m, c.Response); err != nil {
		h.error(c, err)
	}
}
