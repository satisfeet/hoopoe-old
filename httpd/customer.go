package httpd

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/store"
)

type Customer struct {
	store *store.CustomerHandler
}

func NewCustomer(s *store.Store) *Customer {
	h := store.NewCustomerHandler(s)

	return &Customer{h}
}

func (c *Customer) List(context *context.Context) {
	query := store.Query{}
	result := []store.Customer{}

	if p := context.Query("search"); len(p) != 0 {
		query.Search(p, append(store.CustomerIndices, store.CustomerUnique...))
	}

	if err := c.store.FindAll(query, &result); err != nil {
		context.Error(err, 500)

		return
	}
	if err := context.Respond(result, 200); err != nil {
		context.Error(err, 500)
	}
}

func (c *Customer) Show(context *context.Context) {

}

func (c *Customer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := httprouter.New()

	m.Handle("GET", "/customers", context.HandleFunc(c.List))
	m.Handle("GET", "/customers/:id", context.HandleFunc(c.Show))

	m.ServeHTTP(w, r)
}
