package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/httpd/context"
	"github.com/satisfeet/hoopoe/httpd/route"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Customers struct {
	Store *store.Store
}

func (h *Customers) list(c *context.Context) {
	m := []model.Customer{}
	q := store.Query{}

	if s := c.Query("search"); len(s) > 0 {
		if err := q.Search(s, model.CustomerIndex); err != nil {
			c.Error(err, http.StatusBadRequest)

			return
		}
	}

	if err := h.Store.FindAll(q, &m); err != nil {
		c.Error(err, http.StatusInternalServerError)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *Customers) show(c *context.Context) {
	m := model.Customer{}
	q := store.Query{}

	if err := q.Id(c.Param("id")); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.Store.FindOne(q, &m); err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	c.Respond(m, http.StatusOK)
}

func (h *Customers) create(c *context.Context) {
	m := model.Customer{}

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
	m := model.Customer{}

	q := store.Query{}
	q.Id(c.Param("id"))

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.Store.Update(q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customers) destroy(c *context.Context) {
	q := store.Query{}

	if err := q.Id(c.Param("id")); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.Store.Remove(q); err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	c.Respond(nil, http.StatusNoContent)
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
