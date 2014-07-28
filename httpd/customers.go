package httpd

import (
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/hoopoe/store"
)

type Customers struct {
	Store *store.CustomerStore
}

func NewCustomers() *Customers {
	s := store.NewCustomerStore()

	return &Customers{
		Store: s,
	}
}

func (h *Customers) list(c *context.Context) {
	m := []store.Customer{}
	//q.Search(c.Query("search"))

	if err := h.Store.All(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) show(c *context.Context) {
	m := store.Customer{}

	if err := h.Store.One(c.Param("id"), &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) create(c *context.Context) {
	m := store.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.Store.Insert(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) update(c *context.Context) {
	m := store.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.Store.Update(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customers) destroy(c *context.Context) {
	m := store.Customer{}

	if err := h.Store.One(c.Param("id"), &m); err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	if err := h.Store.Remove(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customers) Handler() http.Handler {
	r := router.NewRouter()
	r.HandleFunc(router.Read, "/customers", h.list)
	r.HandleFunc(router.Read, "/customers/:id", h.show)
	r.HandleFunc(router.Create, "/customers", h.create)
	r.HandleFunc(router.Update, "/customers/:id", h.update)
	r.HandleFunc(router.Destroy, "/customers/:id", h.destroy)

	return r
}
