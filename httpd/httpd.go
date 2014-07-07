package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/store"
)

type Httpd struct {
	store *store.Store
}

func New(s *store.Store) *Httpd {
	return &Httpd{s}
}

func (h *Httpd) Listen(addr string) {
	m := http.NewServeMux()

	m.Handle("/", NewCustomer(h.store))

	http.ListenAndServe(addr, Logger(m))
}
