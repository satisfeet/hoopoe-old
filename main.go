package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/go-router"

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
	r := router.NewRouter()

	c, _ := httpd.NewCustomerHandler(s)
	r.Handle("GET", "/customers", httpd.HandlerFunc(c.List))
	r.Handle("POST", "/customers", httpd.HandlerFunc(c.Create))
	r.Handle("GET", "/customers/:customer", httpd.HandlerFunc(c.Show))
	r.Handle("PUT", "/customers/:customer", httpd.HandlerFunc(c.Update))
	r.Handle("DELETE", "/customers/:customer", httpd.HandlerFunc(c.Destroy))

	p, _ := httpd.NewProductHandler(s)
	r.Handle("GET", "/products", httpd.HandlerFunc(p.List))
	r.Handle("POST", "/products", httpd.HandlerFunc(p.Create))
	r.Handle("GET", "/products/:product", httpd.HandlerFunc(p.Show))
	r.Handle("PUT", "/products/:product", httpd.HandlerFunc(p.Update))
	r.Handle("DELETE", "/products/:product", httpd.HandlerFunc(p.Destroy))
	r.Handle("GET", "/products/:product/image", httpd.HandlerFunc(p.Image.Show))
	r.Handle("PUT", "/products/:product/image", httpd.HandlerFunc(p.Image.Update))

	o, _ := httpd.NewOrderHandler(s)
	r.Handle("GET", "/orders", httpd.HandlerFunc(o.List))
	r.Handle("POST", "/orders", httpd.HandlerFunc(o.Create))
	r.Handle("GET", "/orders/:order", httpd.HandlerFunc(o.Show))
	r.Handle("DELETE", "/orders/:order", httpd.HandlerFunc(o.Destroy))

	return handler.Logger(handler.Auth(auth, r))
}
