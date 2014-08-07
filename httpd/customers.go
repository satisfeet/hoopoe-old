package httpd

import (
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/utils"
)

type CustomerHandler struct {
	store  *mgo.Collection
	router *router.Router
}

func NewCustomerHandler(db *mgo.Database) *CustomerHandler {
	r := router.NewRouter()
	c := db.C("customers")
	h := &CustomerHandler{c, r}

	r.HandleFunc(router.Read, "/customers", h.list)
	r.HandleFunc(router.Read, "/customers/:id", h.show)
	r.HandleFunc(router.Create, "/customers", h.create)
	r.HandleFunc(router.Update, "/customers/:id", h.update)
	r.HandleFunc(router.Destroy, "/customers/:id", h.destroy)

	return h
}

func (h *CustomerHandler) list(c *context.Context) {
	q := bson.M{}
	m := []model.Customer{}

	if s := c.Query("search"); len(s) > 0 {
		q["$or"] = make([]bson.M, 0)

		for k, _ := range utils.GetStructInfo(m) {
			m := bson.M{}
			m[k] = bson.RegEx{s, "i"}

			q["$or"] = append(q["$or"].([]bson.M), m)
		}
	}

	if err := h.store.Find(q).All(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *CustomerHandler) show(c *context.Context) {
	q := bson.M{}
	m := model.Customer{}

	if p := c.Param("id"); bson.IsObjectIdHex(p) {
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
	q := bson.M{}

	if p := c.Param("id"); bson.IsObjectIdHex(p) {
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

func (h *CustomerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
