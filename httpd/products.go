package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/model"
)

type ProductHandler struct {
	store  *mgo.Collection
	router *router.Router
}

func NewProductHandler(db *mgo.Database) *ProductHandler {
	r := router.NewRouter()
	c := db.C("products")
	h := &ProductHandler{c, r}

	r.HandleFunc("GET", "/products", h.list)
	r.HandleFunc("GET", "/products/:pid", h.show)
	r.HandleFunc("POST", "/products", h.create)
	r.HandleFunc("PUT", "/products/:pid", h.update)
	r.HandleFunc("DELETE", "/products/:pid", h.destroy)

	return h
}

func (h *ProductHandler) list(c *context.Context) {
	m := []model.Product{}

	if err := h.store.Find(nil).All(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) show(c *context.Context) {
	q := bson.M{}
	m := model.Product{}

	if p := c.Param("pid"); bson.IsObjectIdHex(p) {
		q["_id"] = bson.ObjectIdHex(p)
	} else {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) create(c *context.Context) {
	m := model.Product{Id: bson.NewObjectId()}

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

func (h *ProductHandler) update(c *context.Context) {
	m := model.Product{}

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

func (h *ProductHandler) destroy(c *context.Context) {
	q := bson.M{}

	if p := c.Param("pid"); bson.IsObjectIdHex(p) {
		q["_id"] = bson.ObjectIdHex(p)
	} else {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Remove(q); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
