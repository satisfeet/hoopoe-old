package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store/common"
)

var host, auth, mongodb string

func main() {
	flag.StringVar(&host, "host", ":3000", "")
	flag.StringVar(&auth, "auth", "bodokaiser:secret", "")
	flag.StringVar(&mongodb, "mongo", "localhost/satisfeet", "")
	flag.Parse()

	s := common.NewSession()

	if err := s.Dial(mongodb); err != nil {
		log.Fatalf("Error connecting to database: %s.\n", err)
	}

	if err := http.ListenAndServe(host, Handler(s)); err != nil {
		log.Fatalf("Error starting http server: %s.\n", err)
	}
}

func Handler(s *common.Session) http.Handler {
	r := httpd.NewRouter()

	c, _ := httpd.NewCustomerHandler(s)
	r.HandleFunc("GET", "/customers", c.List)
	r.HandleFunc("POST", "/customers", c.Create)
	r.HandleFunc("GET", "/customers/:customer", c.Show)
	r.HandleFunc("PUT", "/customers/:customer", c.Update)
	r.HandleFunc("DELETE", "/customers/:customer", c.Destroy)

	p, _ := httpd.NewProductHandler(s)
	r.HandleFunc("GET", "/products", p.List)
	r.HandleFunc("POST", "/products", p.Create)
	r.HandleFunc("GET", "/products/:product", p.Show)
	r.HandleFunc("PUT", "/products/:product", p.Update)
	r.HandleFunc("DELETE", "/products/:product", p.Destroy)
	r.HandleFunc("GET", "/products/:product/image", p.Image.Show)
	r.HandleFunc("PUT", "/products/:product/image", p.Image.Update)

	o, _ := httpd.NewOrderHandler(s)
	r.HandleFunc("GET", "/orders", o.List)
	r.HandleFunc("POST", "/orders", o.Create)
	r.HandleFunc("GET", "/orders/:order", o.Show)
	r.HandleFunc("DELETE", "/orders/:order", o.Destroy)

	return handler.Logger(handler.Auth(auth, r))
}
