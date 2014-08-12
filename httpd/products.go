package httpd

import (
	"io"
	"net/http"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type ProductHandler struct {
	store  *store.ProductStore
	router *router.Router
}

func NewProductHandler(s *mongo.Store) *ProductHandler {
	h := &ProductHandler{
		store:  store.NewProductStore(s),
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
	m := []store.Product{}

	if err := h.store.Find(&m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) Show(c *context.Context) {
	m := store.Product{}

	if err := h.store.FindId(c.Param("pid"), &m); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(m, http.StatusOK)
	}
}

func (h *ProductHandler) Create(c *context.Context) {
	m := store.Product{}

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

func (h *ProductHandler) Update(c *context.Context) {
	m := store.Product{}

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

func (h *ProductHandler) Destroy(c *context.Context) {
	if err := h.store.RemoveId(c.Param("pid")); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) ShowImage(c *context.Context) {
	f, err := h.store.OpenImage(c.Param("pid"), c.Param("iid"))
	if err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		io.Copy(c.Response, f)
	}
}

func (h *ProductHandler) CreateImage(c *context.Context) {
	f, err := h.store.CreateImage(c.Param("pid"))
	if err != nil {
		c.Error(err, ErrorCode(err))

		return
	}

	if _, err := io.Copy(f, c.Request.Body); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) DestroyImage(c *context.Context) {
	if err := h.store.RemoveImage(c.Param("pid"), c.Param("iid")); err != nil {
		c.Error(err, ErrorCode(err))
	} else {
		c.Respond(nil, http.StatusNoContent)
	}
}

func (h *ProductHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
