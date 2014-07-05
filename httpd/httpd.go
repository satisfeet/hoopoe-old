package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/store"
)

type Httpd struct {
	store *store.Store
}

func New(store *store.Store) *Httpd {
	return &Httpd{store}
}

func (h *Httpd) Listen(config conf.Map) {
	h.Handle(NewCustomers(h.store))

	http.ListenAndServe(config["addr"], nil)
}

func (h *Httpd) Handle(handler http.Handler) {
	handler = Logger(handler)
	handler = Accept(handler)

	http.Handle("/", handler)
}
