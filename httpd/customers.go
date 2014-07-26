package httpd

import (
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/httpd/route"
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

func (h *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a, p := route.Route("/customers", r)

	c := &context.Context{
		Params: map[string]string{
			"id": p,
		},
		Request:  r,
		Response: w,
	}

	switch a {
	case route.List:
		h.list(c)
	case route.Show:
		h.show(c)
	case route.Create:
		h.create(c)
	case route.Update:
		h.update(c)
	case route.Destroy:
		h.destroy(c)
	default:
		c.Error(nil, http.StatusNotFound)
	}
}
