package main

import (
	"log"
	"net/http"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
)

func main() {
	if err := http.ListenAndServe(host, Handler(s)); err != nil {
		log.Fatalf("Error starting http server: %s.\n", err)
	}
}

func Handler(s *common.Session) http.Handler {
	r := httpd.NewRouter()

	ch := &httpd.CustomerHandler{}

	r.HandleFunc("GET", "/customers", c.List)
	r.HandleFunc("POST", "/customers", c.Create)
	r.HandleFunc("GET", "/customers/:customer", c.Show)
	r.HandleFunc("PUT", "/customers/:customer", c.Update)
	r.HandleFunc("DELETE", "/customers/:customer", c.Destroy)

	return handler.Logger(handler.Auth(auth, r))
}
