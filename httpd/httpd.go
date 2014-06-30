package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/httpd/customers"
	"github.com/satisfeet/hoopoe/httpd/router"
)

func Listen(config map[string]string) {
	r := router.New()

	r.Use(router.Auth)
	r.Use(router.Logger)

	r.Get("/customers", customers.List)
	r.Pos("/customers", customers.Create)
	r.Get("/customers/:customer", customers.Show)
	r.Put("/customers/:customer", customers.Update)
	r.Del("/customers/:customer", customers.Destroy)

	http.ListenAndServe(config["addr"], r)
}
