package httpd

import (
	"net/http"
	"strings"

	"github.com/satisfeet/hoopoe/httpd/context"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type Customers struct {
	Store *store.Store
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
	c := &context.Context{
		Params:   map[string]string{},
		Request:  r,
		Response: w,
	}

	if strings.HasPrefix(r.URL.Path, "/customers") {
		switch s := strings.Split(r.URL.Path, "/"); len(s) {
		case 2:
			switch r.Method {
			case "GET":
				h.list(c)
			case "POST":
				h.create(c)
			}
			return
		case 3:
			switch c.Params["id"] = s[2]; r.Method {
			case "GET":
				h.show(c)
			case "PUT":
				h.update(c)
			case "DELETE":
				h.destroy(c)
			}
			return
		}
	}

	c.Error(nil, http.StatusNotFound)
}
