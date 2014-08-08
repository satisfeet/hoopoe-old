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
	id := store.ParseId(c.Param("pid"))

	m := model.Product{}
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
	id := store.ParseId(c.Param("pid"))

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

func (h *ProductHandler) ShowImage(c *context.Context) {
	pid := store.ParseId(c.Param("pid"))
	iid := store.ParseId(c.Param("iid"))

	q := store.Query{}
	q.Id(pid)

	if !pid.Valid() || !iid.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(nil); err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	file, err := h.files.OpenId(iid)

	if err != nil {
		c.Error(nil, http.StatusNotFound)

		return
	}

	defer file.Close()

	io.Copy(c.Response, file)
}

func (h *ProductHandler) CreateImage(c *context.Context) {
	pid := store.ParseId(c.Param("pid"))
	iid := bson.NewObjectId()

	q := store.Query{}
	q.Id(pid)

	u := store.Update{}
	u.Push("images", iid)

	if !pid.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Find(q).One(nil); err != nil {
		c.Error(err, http.StatusNotFound)

		return
	}

	file, err := h.files.Create("")

	if err != nil {
		c.Error(err, http.StatusInternalServerError)

		return
	}

	file.SetId(iid)
	defer file.Close()

	if _, err := io.Copy(file, c.Request.Body); err != nil {
		c.Error(err, http.StatusBadRequest)

		return
	}

	if _, err := h.store.Upsert(q, u); err != nil {
		c.Error(err, http.StatusInternalServerError)

		return
	}

	c.Respond(nil, http.StatusNoContent)
}

func (h *ProductHandler) DestroyImage(c *context.Context) {
	pid := store.ParseId(c.Param("pid"))
	iid := store.ParseId(c.Param("iid"))

	q := store.Query{}
	q.Id(pid)
	q.In("images", iid)

	u := store.Update{}
	u.Pull("images", iid)

	if !pid.Valid() || !iid.Valid() {
		c.Error(nil, http.StatusBadRequest)

		return
	}

	if err := h.store.Update(q, u); err != nil {
		c.Error(nil, http.StatusNotFound)

		return
	}

	if err := h.files.RemoveId(iid); err != nil {
		c.Error(nil, http.StatusNotFound)
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
