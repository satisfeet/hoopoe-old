package httpd

import (
	"log"
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

// Logs method and url of each incoming request.
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}
