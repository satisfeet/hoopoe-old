package httpd

import (
    "log"
    "net/http"

    "github.com/satisfeet/hoopoe/lib/httpd/mux"
)

func Listen(c map[string]string) {
    m := mux.New()

    m.Use(logger)

    m.Get("/customers", CustomersList)
    m.Pos("/customers", CustomersCreate)
    m.Get("/customers/:customer", CustomersShow)
    m.Put("/customers/:customer", CustomersUpdate)
    m.Del("/customers/:customer", CustomersDestroy)

    http.ListenAndServe(c["port"], m)
}

func logger(c *mux.Context) {
    log.Printf("Request: %s %s", c.Method(), c.Path())

    c.Next()
}
