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
	h := &CustomerHandler{
		store:  db.C("customers"),
		router: router.NewRouter(),
	}

	h.router.HandleFunc("GET", "/customers", h.List)
	h.router.HandleFunc("GET", "/customers/:cid", h.Show)
	h.router.HandleFunc("POST", "/customers", h.Create)
	h.router.HandleFunc("PUT", "/customers/:cid", h.Update)
	h.router.HandleFunc("DELETE", "/customers/:cid", h.Destroy)

	return h
}

func (h *CustomerHandler) List(c *context.Context) {
	m := []model.Customer{}
	q := store.Query{}
	q.Search(c.Query("search"), m)

	if err := h.store.Find(q).All(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) Show(c *context.Context) {
	id := store.ParseId(c.Param("cid"))

	m := model.Customer{}
	q := store.Query{}
	q.Id(id)

	if !id.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) Create(c *context.Context) {
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

func (h *CustomerHandler) Update(c *context.Context) {
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

func (h *CustomerHandler) Destroy(c *context.Context) {
	id := store.ParseId(c.Param("cid"))

	q := store.Query{}
	q.Id(id)

	if !id.Valid() {
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
