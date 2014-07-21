package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/httpd/context"
	"github.com/satisfeet/hoopoe/httpd/router"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Customers struct {
	Store  *store.Store
	router *router.Router
}

func (h *Customers) list(c *context.Context) {
	m := []model.Customer{}

	q := store.Query{}
	q.Search(c.Query("search"), model.CustomerIndex)

	if err := h.Store.FindAll(q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) show(c *context.Context) {
	m := model.Customer{}

	q := store.Query{}
	q.Id(c.Param("id"))

	if err := h.Store.FindOne(q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) create(c *context.Context) {
	m := model.Customer{}

	if c.Parse(&m) {
		if err := h.Store.Insert(&m); err != nil {
			c.Error(err, ErrorCode(err))
		} else {
			c.Respond(m, http.StatusOK)
		}
	}
}

func (h *Customers) update(c *context.Context) {
	m := model.Customer{}

	q := store.Query{}
	q.Id(c.Param("id"))

	if c.Parse(&m) {
		if err := h.Store.Update(q, &m); err != nil {
			c.Error(err, ErrorCode(err))
		} else {
			c.Respond(nil, http.StatusNoContent)
		}
	}
}

func (h *Customers) destroy(c *context.Context) {
	q := store.Query{}
	q.Id(c.Param("id"))

	if err := h.Store.Remove(q); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.router == nil {
		r := router.NewRouter()
		r.HandleFunc(router.MethodShow, "/customers", h.list)
		r.HandleFunc(router.MethodShow, "/customers/:id", h.show)
		r.HandleFunc(router.MethodCreate, "/customers", h.create)
		r.HandleFunc(router.MethodUpdate, "/customers/:id", h.update)
		r.HandleFunc(router.MethodDelete, "/customers/:id", h.destroy)

		h.router = r
	}

	h.router.ServeHTTP(w, r)
}
