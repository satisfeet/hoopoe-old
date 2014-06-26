package httpd

import (
    "log"
    "net/http"

    "github.com/satisfeet/hoopoe/lib/httpd/mux"
)

func Listen(c map[string]string) {
    m := mux.New()

    m.Use(logger)

    m.Get("/", handle)

    http.ListenAndServe(c["port"], m)
}

func logger(c *mux.Context) {
    log.Printf("Request: %s %s", c.Method(), c.Path())

    c.Next()
}

func handle(c *mux.Context) {
    c.Respond("Hello World", 200)
}
