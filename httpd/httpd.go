package httpd

import (
    "net/http"

    "github.com/satisfeet/hoopoe/httpd/router"
)

func Listen(config map[string]string) {
    r := router.New()

    r.Use(router.Auth)
    r.Use(router.Logger)

    r.Get("/customers", CustomersList)
    r.Pos("/customers", CustomersCreate)
    r.Get("/customers/:customer", CustomersShow)
    r.Put("/customers/:customer", CustomersUpdate)
    r.Del("/customers/:customer", CustomersDestroy)

    http.ListenAndServe(config["addr"], r)
}
