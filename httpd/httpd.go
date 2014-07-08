package httpd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/satisfeet/hoopoe/store"
)

type Httpd struct {
	store *store.Store
}

func NewHttpd(s *store.Store) *Httpd {
	return &Httpd{s}
}

func (h *Httpd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m := http.NewServeMux()

	m.Handle("/customers", &CustomerAPI{h.store})

	m.HandleFunc("/", NotFound)

	Logger(m).ServeHTTP(w, r)
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())

		h.ServeHTTP(w, r)
	})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	m := http.StatusText(http.StatusNotFound)

	http.Error(w, fmt.Sprintf("{\"error\":\"%s\"}", m), http.StatusNotFound)
}
