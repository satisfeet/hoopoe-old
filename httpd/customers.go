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

}

func (c *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := httprouter.New()

	m.GET("/customers", c.List)
	m.GET("/customers/:customer", c.Show)

	m.ServeHTTP(w, r)
}
