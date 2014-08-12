package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store"
)

type CustomerHandler struct {
	store  *store.CustomerStore
	router *router.Router
}

func NewCustomerHandler(db *mgo.Database) *CustomerHandler {
	h := &CustomerHandler{
		store:  store.NewCustomerStore(db),
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
	m := []store.Customer{}

	if err := h.store.Search(c.Query("search"), &m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) Show(c *context.Context) {
	m := store.Customer{}
	m.Id = store.ParseId(c.Param("cid"))

	if !m.Id.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.FindId(m.Id, &m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) Create(c *context.Context) {
	m := store.Customer{Id: bson.NewObjectId()}

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
	m := store.Customer{}

	if err := c.Parse(&m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := validation.Validate(m); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if err := h.store.Update(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) Destroy(c *context.Context) {
	m := store.Customer{}
	m.Id = store.ParseId(c.Param("cid"))

	if !m.Id.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Remove(m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
