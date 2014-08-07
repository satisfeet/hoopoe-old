package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type CustomerHandler struct {
	store  *mgo.Collection
	router *router.Router
}

func NewCustomerHandler(db *mgo.Database) *CustomerHandler {
	r := router.NewRouter()
	c := db.C("customers")
	h := &CustomerHandler{c, r}

	r.HandleFunc("GET", "/customers", h.list)
	r.HandleFunc("GET", "/customers/:cid", h.show)
	r.HandleFunc("POST", "/customers", h.create)
	r.HandleFunc("PUT", "/customers/:cid", h.update)
	r.HandleFunc("DELETE", "/customers/:cid", h.destroy)

	return h
}

func (h *CustomerHandler) list(c *context.Context) {
	m := []model.Customer{}

	q := store.Query{}
	q.Search(c.Query("search"), m)

	if err := h.store.Find(q).All(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) show(c *context.Context) {
	m := model.Customer{}

	q := store.Query{}

	if err := q.Id(c.Param("cid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) create(c *context.Context) {
	m := model.Customer{Id: bson.NewObjectId()}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := validation.Validate(m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Insert(m); err != nil {
		c.Error(err, http.StatusInternalServerError)
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

	if err := validation.Validate(m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.UpdateId(m.Id, &m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) destroy(c *context.Context) {
	q := store.Query{}

	if err := q.Id(c.Param("cid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Remove(q); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
