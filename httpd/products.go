package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satisfeet/hoopoe/store"
)

var (
	ProductsPolyPath = "/products"
	ProductsMonoPath = "/products/{id}"
)

type ProductsHandler struct {
	mux   *mux.Router
	store *store.Store
}

func NewProductsHandler(s *store.Store) *ProductsHandler {
	m := mux.NewRouter()

	return &ProductsHandler{
		mux:   m,
		store: s,
	}
}

func (h *ProductsHandler) list(w http.ResponseWriter, r *http.Request) {
	q := store.Query{}
	c := []store.Product{}

	if err := h.store.FindAllProduct(q, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (h *ProductsHandler) show(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	q := store.Query{}
	c := store.Product{}

	if err := q.Id(p["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.FindOneProduct(q, &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (h *ProductsHandler) create(w http.ResponseWriter, r *http.Request) {
	c := store.Product{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.InsertProduct(&c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (h *ProductsHandler) update(w http.ResponseWriter, r *http.Request) {
	c := store.Product{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.UpdateProduct(&c); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductsHandler) destroy(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	q := store.Query{}

	if err := q.Id(p["id"]); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	if err := h.store.RemoveProduct(q); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.Path(ProductsPolyPath).HandlerFunc(h.list).Methods("GET")
	h.mux.Path(ProductsMonoPath).HandlerFunc(h.show).Methods("GET")
	h.mux.Path(ProductsPolyPath).HandlerFunc(h.create).Methods("POST")
	h.mux.Path(ProductsMonoPath).HandlerFunc(h.update).Methods("PUT")
	h.mux.Path(ProductsMonoPath).HandlerFunc(h.destroy).Methods("DELETE")

	h.mux.ServeHTTP(w, r)
}
