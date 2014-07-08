package httpd

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/store"
)

type Customer struct {
	store *store.Store
}

func NewCustomer(s *store.Store) *Customer {
	return &Customer{s}
}

func (c *Customer) List(context *context.Context) {
	query := store.Query{}
	result := []store.Customer{}

	if key := context.Query("search"); len(key) != 0 {
		query.Search(key, append(store.CustomerIndices, store.CustomerUnique...))
	}
	if err := store.FindAllCustomer(c.store, query, &result); err != nil {
		context.Error(err, 500)

		return
	}
	if err := context.Respond(result, 200); err != nil {
		context.Error(err, 500)
	}
}

func (c *Customer) Show(context *context.Context) {
	query := store.Query{}
	result := store.Customer{}

	if err := query.IdHex(context.Param("id")); err != nil {
		context.Error(err, 404)

		return
	}
	if err := store.FindOneCustomer(c.store, query, &result); err != nil {
		context.Error(err, 500)

		return
	}
	if err := context.Respond(result, 200); err != nil {
		context.Error(err, 500)
	}
}

func (c *Customer) Create(context *context.Context) {
	result := store.Customer{}

	if err := context.Parse(&result); err != nil {
		context.Error(err, 400)

		return
	}
	if err := store.InsertCustomer(c.store, &result); err != nil {
		context.Error(err, 500)

		return
	}
	if err := context.Respond(result, 200); err != nil {
		context.Error(err, 500)
	}
}

func (c *Customer) Update(context *context.Context) {
	query := store.Query{}
	result := store.Customer{}

	if err := query.IdHex(context.Param("id")); err != nil {
		context.Error(err, 404)

		return
	}
	if err := context.Parse(&result); err != nil {
		context.Error(err, 400)

		return
	}
	if err := store.UpdateCustomer(c.store, &result); err != nil {
		context.Error(err, 500)

		return
	}
	if err := context.Respond(nil, 204); err != nil {
		context.Error(err, 500)
	}
}

func (c *Customer) Destroy(context *context.Context) {
	query := store.Query{}

	if err := query.IdHex(context.Param("id")); err != nil {
		context.Error(err, 404)

		return
	}
	if err := store.RemoveCustomer(c.store, query); err != nil {
		context.Error(err, 500)

		return
	}
	if err := context.Respond(nil, 204); err != nil {
		context.Error(err, 500)
	}
}

func (c *Customer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := httprouter.New()

	m.Handle("POST", "/customers", context.HandleFunc(c.Create))
	m.Handle("GET", "/customers", context.HandleFunc(c.List))
	m.Handle("GET", "/customers/:id", context.HandleFunc(c.Show))
	m.Handle("PUT", "/customers/:id", context.HandleFunc(c.Update))
	m.Handle("DELETE", "/customers/:id", context.HandleFunc(c.Destroy))

	m.ServeHTTP(w, r)
}
