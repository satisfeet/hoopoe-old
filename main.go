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

	if err := c.Flags(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing arguments: %s.\n", err)

		return
	}
	if err := store.Open(c.Mongo); err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	h := httpd.Logger(&httpd.Auth{
		Username: c.Username,
		Password: c.Password,
		Handler:  Handler(),
	})

	if err := http.ListenAndServe(c.Host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}

// Handler sets up a basic prefixed based http multiplexer to switch between
// different HTTP resources. If not HTTP resource is found it will respond a Not
// Found error.
func Handler() http.Handler {
	c := httpd.NewCustomers(&store.Store{
		Name: "customers",
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
