package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

func main() {
	c := &conf.Conf{}
	s := &store.Session{}

	if err := c.Flags(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing arguments: %s.\n", err)

		return
	}
	if err := s.Open(c.Mongo); err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	httpd.Basic = c.Auth

	h := httpd.Logger(httpd.Auth(Handler(s)))

	if err := http.ListenAndServe(c.Host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}

func Handler(s *store.Session) http.Handler {
	c := httpd.NewCustomers(&store.Store{
		Name:    "customers",
		Session: s,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/customers"):
			c.ServeHTTP(w, r)
		default:
			httpd.NotFound(w, r)
		}
	})
}
