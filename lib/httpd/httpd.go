package httpd

import (
    "log"
    "errors"
    "strings"
    "net/http"
    "encoding/base64"

    "github.com/satisfeet/hoopoe/lib/httpd/mux"
)

func Listen(c map[string]string) {
    m := mux.New()

    m.Use(auth)
    m.Use(logger)

    m.Get("/customers", CustomersList)
    m.Pos("/customers", CustomersCreate)
    m.Get("/customers/:customer", CustomersShow)
    m.Put("/customers/:customer", CustomersUpdate)
    m.Del("/customers/:customer", CustomersDestroy)

    http.ListenAndServe(c["port"], m)
}

func auth(c *mux.Context) {
    h := strings.Split(c.Get("Authorization"), " ")

    if len(h) != 2 {
        c.RespondError(errors.New("Unauthorized"), 401)

        return
    }

    b, err := base64.StdEncoding.DecodeString(h[1])

    if err != nil {
        c.RespondError(errors.New("Bad Request"), 400)

        return
    }

    p := strings.Split(string(b), ":")

    if p[0] != "bodokaiser" {
        c.RespondError(errors.New("Unauthorized"), 401)

        return
    }

    c.Next()
}

func logger(c *mux.Context) {
    log.Printf("Request: %s %s", c.Method(), c.Path())

    c.Next()
}
