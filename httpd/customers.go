package httpd

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/satisfeet/hoopoe/store"
)

type Customers struct {
	manager *store.Manager
}

func NewCustomers(s *store.Store) *Customers {
	m := s.Manager("customers")

	return &Customers{m}
}

func (c *Customers) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := []store.Customer{}

	q := store.Query{}

	if param := r.URL.Query().Get("search"); len(param) != 0 {
		q.Search(param)
	}

	if err := c.manager.Find(q, &result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (c *Customers) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := store.Customer{}

	q := store.Query{}
	q.Id(p.ByName("customer"))

	if err := c.manager.FindOne(q, &result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (c *Customers) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := store.Customer{}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := c.manager.Create(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := json.NewEncoder(w).Encode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (c *Customers) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := store.Customer{}

	q := store.Query{}
	q.Id(p.ByName("customer"))

	if err := c.manager.FindOne(q, &result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := c.manager.Update(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Customers) Destroy(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result := store.Customer{}

	q := store.Query{}
	q.Id(p.ByName("customer"))

	if err := c.manager.FindOne(q, &result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if err := c.manager.Destroy(&result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := httprouter.New()

	m.GET("/customers", c.List)
	m.POST("/customers", c.Create)
	m.GET("/customers/:customer", c.Show)
	m.PUT("/customers/:customer", c.Update)
	m.DELETE("/customers/:customer", c.Destroy)

	m.ServeHTTP(w, r)
}
