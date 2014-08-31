package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

var Url = os.Getenv("HOOPOE_MYSQL")
var Host = os.Getenv("HOOPOE_HOST")
var Auth = os.Getenv("HOOPOE_AUTH")

func main() {
	db, err := sql.Open("mysql", Url)

	if err != nil {
		log.Fatalf("Error connecting to db: %s.\n", err)
	}

	if err := http.ListenAndServe(Host, Handler(db)); err != nil {
		log.Fatalf("Error starting http server: %s.\n", err)
	}
}

func Handler(db *sql.DB) http.Handler {
	r := httpd.NewRouter()

	c := &httpd.CustomerHandler{
		Store: store.NewCustomerStore(db),
	}
	p := &httpd.ProductHandler{
		Store: store.NewProductStore(db),
	}

	r.HandleFunc("GET", "/products", p.List)
	r.HandleFunc("GET", "/customers", c.List)

	r.HandleFunc("GET", "/products/:product", p.Show)
	r.HandleFunc("GET", "/customers/:customer", c.Show)

	r.HandleFunc("POST", "/customers", c.Create)

	return handler.Logger(handler.Auth(Auth, r))
}
