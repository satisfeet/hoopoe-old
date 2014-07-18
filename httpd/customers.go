package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type CustomersHandler struct {
	store  *store.Store
	router *httprouter.Router
}

func NewCustomersHandler(n string) *CustomersHandler {
	return &CustomersHandler{
		store:  store.NewStore(n),
		router: httprouter.New(),
	}
}

func (h *CustomersHandler) list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := store.Query{}
	m := []model.Customer{}

	if s := r.URL.Query().Get("search"); len(s) != 0 {
		q.Search(s, model.CustomerIndex)
	}
	if err := h.store.FindAll(q, &m); err != nil {
		Error(w, err, http.StatusInternalServerError)

		return
	}

	Respond(w, m, http.StatusOK)
}

func (h *CustomersHandler) show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (h *CustomersHandler) create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *CustomersHandler) update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (h *CustomersHandler) destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (h *CustomersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.Handle("POST", "/customers", h.create)
	h.router.Handle("GET", "/customers", h.list)
	h.router.Handle("GET", "/customers/:id", h.show)
	h.router.Handle("PUT", "/customers/:id", h.update)
	h.router.Handle("DELETE", "/customers/:id", h.destroy)

	h.router.ServeHTTP(w, r)
}
