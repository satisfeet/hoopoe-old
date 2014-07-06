package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/store"
)

type Httpd struct {
	store *store.Store
}

func New(s *store.Store) *Httpd {
	return &Httpd{s}
}

func (h *Httpd) Listen(c conf.Map) {
	m := http.NewServeMux()

	m.Handle("/", NewCustomer(h.store))

	http.ListenAndServe(c["addr"], Logger(m))
}
