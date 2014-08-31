package main

import (
	"log"
	"net/http"
	"os"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
	"github.com/satisfeet/hoopoe/store/common"
)

var Url = os.Getenv("HOOPOE_MYSQL")
var Host = os.Getenv("HOOPOE_HOST")
var Auth = os.Getenv("HOOPOE_AUTH")

func main() {
	s, err := store.Open(Url)

	if err != nil {
		log.Fatalf("Error connecting to db: %s.\n", err)
	}

	if err := http.ListenAndServe(Host, Handler(s)); err != nil {
		log.Fatalf("Error starting http server: %s.\n", err)
	}
}

func Handler(s *common.Session) http.Handler {
	r := httpd.NewRouter()

	c := &httpd.CustomerHandler{
		Store: store.NewCustomerStore(s),
	}
	p := &httpd.ProductHandler{
		Store: store.NewProductStore(s),
	}

	r.HandleFunc("GET", "/customers", c.List)
	r.HandleFunc("GET", "/products", p.List)

	r.HandleFunc("GET", "/customers/:customer", c.Show)

	return handler.Logger(handler.Auth(Auth, r))
}
