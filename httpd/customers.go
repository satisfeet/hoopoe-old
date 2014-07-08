package httpd

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/store"
)

type CustomerAPI struct {
	store *store.Store
}

func (ca *CustomerAPI) list(c *context.Context) {
	query := store.Query{}
	result := []store.Customer{}

	if key := c.Query("search"); len(key) != 0 {
		query.Search(key, append(store.CustomerIndices, store.CustomerUnique...))
	}

	if err := store.FindAllCustomer(ca.store, query, &result); err == nil {
		if err := c.Respond(result, 200); err != nil {
			c.Error(err, 500)
		}
	} else {
		c.Error(err, 500)
	}
}

func (ca *CustomerAPI) show(c *context.Context) {
	query := store.Query{}
	result := store.Customer{}

	if err := query.Id(c.Param("id")); err != nil {
		c.Error(err, 404)

		return
	}

	if err := store.FindOneCustomer(ca.store, query, &result); err == nil {
		if err := c.Respond(result, 200); err != nil {
			c.Error(err, 500)
		}
	} else {
		c.Error(err, 500)
	}
}

func (ca *CustomerAPI) create(c *context.Context) {
	result := store.Customer{}

	if err := c.Parse(&result); err != nil {
		c.Error(err, 400)

		return
	}

	if err := store.InsertCustomer(ca.store, &result); err == nil {
		if err := c.Respond(result, 200); err != nil {
			c.Error(err, 500)
		}
	} else {
		c.Error(err, 500)
	}
}

func (ca *CustomerAPI) update(c *context.Context) {
	result := store.Customer{}

	if err := c.Parse(&result); err != nil {
		c.Error(err, 400)

		return
	}

	if err := store.UpdateCustomer(ca.store, &result); err == nil {
		if err := c.Respond(nil, 204); err != nil {
			c.Error(err, 500)
		}
	} else {
		c.Error(err, 500)
	}
}

func (ca *CustomerAPI) destroy(c *context.Context) {
	query := store.Query{}

	if err := query.Id(c.Param("id")); err != nil {
		c.Error(err, 404)

		return
	}

	if err := store.RemoveCustomer(ca.store, query); err == nil {
		if err := c.Respond(nil, 204); err != nil {
			c.Error(err, 500)
		}
	} else {
		c.Error(err, 404)
	}
}

func (ca *CustomerAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := httprouter.New()

	m.Handle("GET", "/customers", context.HandleFunc(ca.list))
	m.Handle("POST", "/customers", context.HandleFunc(ca.create))
	m.Handle("GET", "/customers/:id", context.HandleFunc(ca.show))
	m.Handle("PUT", "/customers/:id", context.HandleFunc(ca.update))
	m.Handle("DELETE", "/customers/:id", context.HandleFunc(ca.destroy))

	m.ServeHTTP(w, r)
}
