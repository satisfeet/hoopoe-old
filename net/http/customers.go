package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satisfeet/hoopoe/store"
)

var (
	CustomersPolyPath = "/customers"
	CustomersMonoPath = "/customers/{id}"
)

type CustomersHandler struct {
	mux   *mux.Router
	store *store.Store
}

func NewCustomersHandler(s *store.Store) *CustomersHandler {
	m := mux.NewRouter()

	return &CustomersHandler{
		mux:   m,
		store: s,
	}
}

func (h *CustomersHandler) list(w http.ResponseWriter, r *http.Request) {
	q := store.Query{}
	c := []store.Customer{}

	if s := r.URL.Query().Get("search"); len(s) != 0 {
		q.Search(s, append(
			store.CustomerIndices,
			store.CustomerUnique...,
		))
	}
	if err := h.store.FindAllCustomer(q, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (h *CustomersHandler) show(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	q := store.Query{}
	c := store.Customer{}

	if err := q.Id(p["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.FindOneCustomer(q, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (h *CustomersHandler) create(w http.ResponseWriter, r *http.Request) {
	c := store.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.InsertCustomer(&c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (h *CustomersHandler) update(w http.ResponseWriter, r *http.Request) {
	c := store.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.UpdateCustomer(&c); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CustomersHandler) destroy(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	q := store.Query{}

	if err := q.Id(p["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.RemoveCustomer(q); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CustomersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.Path(CustomersPolyPath).HandlerFunc(h.list).Methods("GET")
	h.mux.Path(CustomersMonoPath).HandlerFunc(h.show).Methods("GET")
	h.mux.Path(CustomersPolyPath).HandlerFunc(h.create).Methods("POST")
	h.mux.Path(CustomersMonoPath).HandlerFunc(h.update).Methods("PUT")
	h.mux.Path(CustomersMonoPath).HandlerFunc(h.destroy).Methods("DELETE")

	h.mux.ServeHTTP(w, r)
}
