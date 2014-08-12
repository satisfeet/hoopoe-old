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
	store  *store.Product
	router *router.Router
}

func NewProductHandler(db *mgo.Database) *ProductHandler {
	h := &ProductHandler{
		store:  store.NewProduct(db),
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

	if err := h.store.Find(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) Show(c *context.Context) {
	m := model.Product{}
	m.Id = store.ParseId(c.Param("pid"))

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

	if err := h.store.Update(&m); err != nil {
		c.Error(err, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) Destroy(c *context.Context) {
	m := model.Product{}
	m.Id = store.ParseId(c.Param("pid"))

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

func (h *ProductHandler) ShowImage(c *context.Context) {
	m := model.Product{}
	m.Id = store.ParseId(c.Param("pid"))

	iid := store.ParseId(c.Param("iid"))

	if !m.Id.Valid() || !iid.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	f, err := h.store.OpenImage(m, store.ParseId(c.Param("iid")))
	if err != nil {
		c.Error(nil, http.StatusNotFound)
	} else {
		io.Copy(c.Response, f)
	}
}

func (h *ProductHandler) CreateImage(c *context.Context) {
	m := model.Product{}
	m.Id = store.ParseId(c.Param("pid"))

	if !m.Id.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	f, err := h.store.CreateImage(m)
	if err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	if _, err := io.Copy(f, c.Request.Body); err != nil {
		c.Error(err, http.StatusBadRequest)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) DestroyImage(c *context.Context) {
	m := model.Product{}
	m.Id = store.ParseId(c.Param("pid"))
	iid := store.ParseId(c.Param("iid"))

	if !m.Id.Valid() || !iid.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.RemoveImage(m, iid); err != nil {
		c.Error(nil, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
