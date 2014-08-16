package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/common"
)

type OrderHandler struct {
	store *store.OrderStore
}

func NewOrderHandler(s *common.Session) (*OrderHandler, error) {
	store, err := store.NewOrderStore(s)

	return &OrderHandler{
		store: store,
	}, err
}

func (h *OrderHandler) List(c *Context) {
	m := []store.Order{}

	if err := h.store.FindAll(nil, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *OrderHandler) Show(c *Context) {
	m := store.Order{}

	q := store.NewOrderQuery()
	q.Id(c.Param("order"))

	if err := h.store.FindOne(q, &m); err != nil {
		c.Error(err)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *OrderHandler) Create(c *Context) {
	m := store.Order{}

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

func (h *OrderHandler) Destroy(c *Context) {
	q := store.NewOrderQuery()
	q.Id(c.Param("order"))

	if err := h.store.Remove(q); err != nil {
		c.Error(err)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}
