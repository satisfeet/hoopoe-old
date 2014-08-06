package httpd

import (
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
	"gopkg.in/mgo.v2/bson"
)

type CustomerHandler struct {
	store  *store.Store
	router *router.Router
}

func NewCustomerHandler(s *store.Store) *CustomerHandler {
	r := router.NewRouter()

	h := &CustomerHandler{
		store:  s,
		router: r,
	}

	r.HandleFunc(router.Read, "/customers", h.list)
	r.HandleFunc(router.Read, "/customers/:id", h.show)
	r.HandleFunc(router.Create, "/customers", h.create)
	r.HandleFunc(router.Update, "/customers/:id", h.update)
	r.HandleFunc(router.Destroy, "/customers/:id", h.destroy)

	return h
}

func (h *CustomerHandler) list(c *context.Context) {
	m := []model.Customer{}

	var err error
	if s := c.Query("search"); len(s) != 0 {
		err = h.store.Search(s, &m)
	} else {
		err = h.store.FindAll(&m)
	}

	if err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) show(c *context.Context) {
	m := model.Customer{}

	if err := h.store.FindId(c.Param("id"), &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) create(c *context.Context) {
	m := model.Customer{
		Id: bson.NewObjectId(),
	}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Insert(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) update(c *context.Context) {
	m := model.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Update(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) destroy(c *context.Context) {
	m := model.Customer{}

	if err := h.store.FindId(c.Param("id"), &m); err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	if err := h.store.Remove(m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
