package httpd

import (
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/satisfeet/hoopoe/store"
)

type Customers struct{}

func (c *Customers) List(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(w, r)

	res, err := CustomersFindAll(Query{
		"search": r.URL.Query().Get("search"),
	})

	if err != nil {
		h.Error(err, 500)
	} else {
		h.Respond(res, 200)
	}
}

func (c *Customers) Show(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(w, r)

	res, err := CustomersFindOne(Query{
		"id": mux.Vars(r)["customer"],
	})

	if err != nil {
		h.Error(err, 500)
	} else {
		h.Respond(res, 200)
	}
}

func (c *Customers) Create(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(w, r)

	res := Customer{}

	if err := h.Parse(&res); err != nil {
		h.Error(err, 500)

		return
	}

	if err := CustomersCreate(&res); err != nil {
		h.Error(err, 500)
	} else {
		h.Respond(res, 200)
	}
}

func (c *Customers) Update(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(w, r)

	res, err := CustomersFindOne(Query{
		"id": mux.Vars(r)["customer"],
	})

	if err != nil {
		h.Error(err, 500)

		return
	}

	if err := h.Parse(&res); err != nil {
		h.Error(err, 500)

		return
	}

	if err := CustomersUpdate(&res); err != nil {
		h.Error(err, 500)
	} else {
		h.Respond(nil, 204)
	}
}

func (c *Customers) Destroy(w http.ResponseWriter, r *http.Request) {
	h := NewHandler(w, r)

	res, err := CustomersFindOne(Query{
		"id": mux.Vars(r)["customer"],
	})

	if err != nil {
		h.Error(err, 500)

		return
	}

	if err := CustomersRemove(&res); err != nil {
		h.Error(err, 500)
	} else {
		h.Respond(nil, 204)
	}
}

func (c *Customers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := mux.NewRouter()

	m.HandleFunc("/customers", c.List).Methods("GET")
	m.HandleFunc("/customers", c.Create).Methods("POST")
	m.HandleFunc("/customers/{customer}", c.Show).Methods("GET")
	m.HandleFunc("/customers/{customer}", c.Update).Methods("PUT")
	m.HandleFunc("/customers/{customer}", c.Destroy).Methods("DELETE")

	m.ServeHTTP(w, r)
}
