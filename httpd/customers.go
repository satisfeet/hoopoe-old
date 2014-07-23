package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/httpd/context"
	"github.com/satisfeet/hoopoe/httpd/router"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Customers struct {
	store  *store.Store
	router *router.Router
}

func NewCustomers(s *store.Store) *Customers {
	r := router.NewRouter()

	c := &Customers{
		store:  s,
		router: r,
	}

	r.HandleFunc(router.MethodShow, "/customers", c.List)
	r.HandleFunc(router.MethodShow, "/customers/:id", c.Show)
	r.HandleFunc(router.MethodCreate, "/customers", c.Create)
	r.HandleFunc(router.MethodUpdate, "/customers/:id", c.Update)
	r.HandleFunc(router.MethodDelete, "/customers/:id", c.Destroy)

	return c
}

func (h *Customers) List(c *context.Context) {
	m := []model.Customer{}

	q := store.Query{}
	q.Search(c.Query("search"), model.CustomerIndex)

	if err := h.store.FindAll(q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) Show(c *context.Context) {
	m := model.Customer{}

	q := store.Query{}
	q.Id(c.Param("id"))

	if err := h.store.FindOne(q, &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *Customers) Create(c *context.Context) {
	m := model.Customer{}

	if c.Parse(&m) {
		if err := h.store.Insert(&m); err != nil {
			c.Error(err, ErrorCode(err))
		} else {
			c.Respond(m, http.StatusOK)
		}
	}
}

func (h *Customers) Update(c *context.Context) {
	m := model.Customer{}

	q := store.Query{}
	q.Id(c.Param("id"))

	if c.Parse(&m) {
		if err := h.store.Update(q, &m); err != nil {
			c.Error(err, ErrorCode(err))
		} else {
			c.Respond(nil, http.StatusNoContent)
		}
	}
}

func (h *Customers) Destroy(c *context.Context) {
	q := store.Query{}
	q.Id(c.Param("id"))

	if err := h.store.Remove(q); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
