package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
)

var host, auth, mongo string

func main() {
	flag.StringVar(&host, "host", ":3000", "")
	flag.StringVar(&auth, "auth", "bodokaiser:secret", "")
	flag.StringVar(&mongo, "mongo", "localhost/satisfeet", "")
	flag.Parse()

	s, err := mgo.Dial(mongo)
	if err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	if err := http.ListenAndServe(host, handle(s.DB(""))); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}

func handle(db *mgo.Database) http.Handler {
	p := httpd.NewProductHandler(db)
	c := httpd.NewCustomerHandler(db)

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix("/products", r.URL.Path):
			p.ServeHTTP(w, r)
		case strings.HasPrefix("/customers", r.URL.Path):
			c.ServeHTTP(w, r)
		}
	})

	return handler.Logger(handler.Auth(auth, h))
}
