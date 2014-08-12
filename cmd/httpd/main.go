package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store/mongo"
)

var host, auth, mongodb string

func main() {
	flag.StringVar(&host, "host", ":3000", "")
	flag.StringVar(&auth, "auth", "bodokaiser:secret", "")
	flag.StringVar(&mongodb, "mongo", "localhost/satisfeet", "")
	flag.Parse()

	s := &mongo.Store{}
	if err := s.Dial(mongodb); err != nil {
		log.Printf("Error connecting to database: %s.\n", err)

		return
	}

	if err := http.ListenAndServe(host, Handle(s)); err != nil {
		log.Printf("Error starting http server: %s.\n", err)
	}
}

func Handle(s *mongo.Store) http.Handler {
	p := httpd.NewProductHandler(s)
	c := httpd.NewCustomerHandler(s)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/products"):
			p.ServeHTTP(w, r)
		case strings.HasPrefix(r.URL.Path, "/customers"):
			c.ServeHTTP(w, r)
		}
	})

	return handler.Logger(handler.Auth(auth, h))
}
