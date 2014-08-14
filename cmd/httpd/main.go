package main

import (
	"flag"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/go-router"
	"github.com/satisfeet/hoopoe/httpd"
)

var host, auth, mongodb string

func main() {
	flag.StringVar(&host, "host", ":3000", "")
	flag.StringVar(&auth, "auth", "bodokaiser:secret", "")
	flag.StringVar(&mongodb, "mongo", "localhost/satisfeet", "")
	flag.Parse()

	s, err := mgo.Dial(mongodb)

	if err != nil {
		log.Printf("Error connecting to database: %s.\n", err)

		return
	}

	if err := http.ListenAndServe(host, Handler(s)); err != nil {
		log.Printf("Error starting http server: %s.\n", err)
	}
}

func Handler(s *mgo.Session) http.Handler {
	p := httpd.NewProduct(s)
	c := httpd.NewCustomer(s)

	r := router.NewRouter()

	r.HandleFunc("GET", "/customers", c.List)
	r.HandleFunc("GET", "/products", p.List)

	r.HandleFunc("POST", "/customers", c.Create)
	r.HandleFunc("POST", "/products", p.Create)

	r.HandleFunc("GET", "/customers/:customer", c.Show)
	r.HandleFunc("GET", "/products/:product", p.Show)

	r.HandleFunc("PUT", "/customers/:customer", c.Update)
	r.HandleFunc("PUT", "/products/:product", p.Update)

	r.HandleFunc("DELETE", "/customers/:customer", c.Destroy)
	r.HandleFunc("DELETE", "/products/:product", p.Destroy)

	r.HandleFunc("POST", "/products/:product/images", p.CreateImage)
	r.HandleFunc("GET", "/products/:product/images/:image", p.ShowImage)
	r.HandleFunc("DELETE", "/products/:product/images/:image", p.DestroyImage)

	return handler.Logger(handler.Auth(auth, r))
}
