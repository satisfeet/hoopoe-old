package httpd

import (
	"io"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type ProductHandler struct {
	files  *mgo.GridFS
	store  *mgo.Collection
	router *router.Router
}

func NewProductHandler(db *mgo.Database) *ProductHandler {
	h := &ProductHandler{
		store:  db.C("products"),
		files:  db.GridFS("products"),
		router: router.NewRouter(),
	}

	h.router.HandleFunc("GET", "/products", h.List)
	h.router.HandleFunc("GET", "/products/:pid", h.Show)
	h.router.HandleFunc("GET", "/products/:pid/images/:iid", h.ShowImage)
	h.router.HandleFunc("POST", "/products", h.Create)
	h.router.HandleFunc("POST", "/products/:pid/images", h.CreateImage)
	h.router.HandleFunc("PUT", "/products/:pid", h.Update)
	h.router.HandleFunc("DELETE", "/products/:pid", h.Destroy)
	h.router.HandleFunc("DELETE", "/products/:pid/images/:iid", h.DestroyImage)

	return h
}

func (h *ProductHandler) List(c *context.Context) {
	m := []model.Product{}

	if err := h.store.Find(nil).All(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) Show(c *context.Context) {
	m := model.Product{}
	q := store.Query{}

	if err := q.Id(c.Param("pid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) Create(c *context.Context) {
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

func (h *ProductHandler) Update(c *context.Context) {
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

func (h *ProductHandler) Destroy(c *context.Context) {
	q := store.Query{}

	if err := q.Id(c.Param("pid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Remove(q); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) ShowImage(c *context.Context) {
	m := model.Product{}
	q := store.Query{}

	if err := q.Id(c.Param("pid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	if p := c.Param("iid"); !bson.IsObjectIdHex(p) {
		c.Error(nil, http.StatusBadRequest)
	} else {
		file, err := h.files.OpenId(bson.ObjectIdHex(p))

		if err != nil {
			c.Error(nil, http.StatusNotFound)

			return
		}

		io.Copy(c.Response, file)
		file.Close()
	}
}

func (h *ProductHandler) CreateImage(c *context.Context) {
	m := model.Product{}
	q := store.Query{}

	if err := q.Id(c.Param("pid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	id := bson.NewObjectId()

	file, err := h.files.Create("")
	file.SetId(id)

	if err != nil {
		c.Error(err, http.StatusInternalServerError)

		return
	}

	defer file.Close()

	if _, err := io.Copy(file, c.Request.Body); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	m.Images = append(m.Images, id)

	if _, err := h.store.UpsertId(m.Id, bson.M{"$push": bson.M{"images": id}}); err != nil {
		c.Error(err, http.StatusInternalServerError)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}

func (h *ProductHandler) DestroyImage(c *context.Context) {
	m := model.Product{}
	q := store.Query{}

	if err := q.Id(c.Param("pid")); err != nil {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(&m); err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	if p := c.Param("iid"); !bson.IsObjectIdHex(p) {
		c.Error(nil, http.StatusBadRequest)
	} else {
		if err := h.files.RemoveId(bson.ObjectIdHex(p)); err != nil {
			c.Error(nil, http.StatusNotFound)
		} else {
			c.Respond(nil, http.StatusNoContent)
		}
	}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
