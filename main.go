package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

func main() {
	c := &conf.Conf{}
	s := &store.Session{}

	if err := c.Flags(os.Args[1:]); err != nil {
		fmt.Errorf("Error parsing arguments: %s.\n", err)

		return
	}
	if err := s.Open(c.Mongo); err != nil {
		fmt.Errorf("Error connecting to database: %s.\n", err)

		return
	}

	http.Handle("/customers", httpd.Auth(&httpd.Customers{
		Store: &store.Store{Name: "customers", Session: s},
	}))
	http.Handle("/", httpd.Auth(httpd.NotFound()))

	if err := http.ListenAndServe(c.Host, nil); err != nil {
		fmt.Errorf("Error starting http server: %s.\n", err)
	}
}
