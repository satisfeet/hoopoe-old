package httpd

import (
	"net/http"

	"github.com/satisfeet/hoopoe/httpd/router"
)

func Init(c map[string]string) {
	r := router.New()

	r.Use(router.Auth)
	r.Use(router.Logger)

	r.Get("/customers", customersList)
	r.Pos("/customers", customersCreate)
	r.Get("/customers/:customer", customersShow)
	r.Put("/customers/:customer", customersUpdate)
	r.Del("/customers/:customer", customersDestroy)

	http.ListenAndServe(c["addr"], r)
}
