package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type ProductsHandler struct {
	store  *store.Store
	router *httprouter.Router
}

func NewProductsHandler(n string) *ProductsHandler {
	return &ProductsHandler{
		store:  store.NewStore(n),
		router: httprouter.New(),
	}
}

func (h *ProductsHandler) list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := store.Query{}
	m := []model.Customer{}

	if err := h.store.FindAll(q, &m); err != nil {
		Error(w, err, http.StatusInternalServerError)

		return
	}

	Respond(w, m, http.StatusOK)
}

func (h *ProductsHandler) show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := store.Query{}
	m := model.Customer{}

	if err := q.Id(p.ByName("id")); err != nil {
		Error(w, err, http.StatusBadRequest)

		return
	}
	if err := h.store.FindOne(q, &m); err != nil {
		Error(w, err, http.StatusInternalServerError)

		return
	}

	Respond(w, m, http.StatusOK)
}

func (h *ProductsHandler) create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	m := model.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		Error(w, err, http.StatusNotFound)

		return
	}
	if err := h.store.Insert(&m); err != nil {
		Error(w, err, http.StatusNotFound)

		return
	}

	Respond(w, m, http.StatusOK)
}

func (h *ProductsHandler) update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := store.Query{}
	m := model.Customer{}

	if err := q.Id(p.ByName("id")); err != nil {
		Error(w, err, http.StatusBadRequest)

		return
	}
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		Error(w, err, http.StatusNotFound)

		return
	}
	if err := h.store.Update(q, &m); err != nil {
		Error(w, err, http.StatusNotFound)

		return
	}

	Respond(w, nil, http.StatusNoContent)
}

func (h *ProductsHandler) destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	q := store.Query{}

	if err := q.Id(p.ByName("id")); err != nil {
		Error(w, err, http.StatusBadRequest)

		return
	}
	if err := h.store.Remove(q); err != nil {
		Error(w, err, http.StatusNotFound)

		return
	}

	Respond(w, nil, http.StatusNoContent)
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.Handle("POST", "/products", h.create)
	h.router.Handle("GET", "/products", h.list)
	h.router.Handle("GET", "/products/:id", h.show)
	h.router.Handle("PUT", "/products/:id", h.update)
	h.router.Handle("DELETE", "/products/:id", h.destroy)

	h.router.ServeHTTP(w, r)
}
