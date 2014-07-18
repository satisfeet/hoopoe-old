package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

type OrdersHandler struct {
	store  *store.Store
	router *httprouter.Router
}

func NewOrdersHandler() *OrdersHandler {
	return &OrdersHandler{
		store:  store.NewStore(model.OrderName),
		router: httprouter.New(),
	}
}

func (h *OrdersHandler) list(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	q := store.Query{}
	m := []model.Customer{}

	if err := h.store.FindAll(q, &m); err != nil {
		Error(w, err, http.StatusInternalServerError)

		return
	}

	Respond(w, m, http.StatusOK)
}

func (h *OrdersHandler) show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (h *OrdersHandler) create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

func (h *OrdersHandler) update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (h *OrdersHandler) destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.Handle("POST", "/orders", h.create)
	h.router.Handle("GET", "/orders", h.list)
	h.router.Handle("GET", "/orders/:id", h.show)
	h.router.Handle("PUT", "/orders/:id", h.update)
	h.router.Handle("DELETE", "/orders/:id", h.destroy)

	h.router.ServeHTTP(w, r)
}
