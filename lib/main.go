package main

import (
    "log"

    "github.com/satisfeet/hoopoe/lib/conf"
    "github.com/satisfeet/hoopoe/lib/store"
    "github.com/satisfeet/hoopoe/lib/httpd"
)

func main() {
    c := conf.New()

    if err := c.FromFlags(); err != nil {
        log.Fatal(err)
    }

    if err := store.Open(c.Store); err != nil {
        log.Fatal(err)
    }

    httpd.Listen(c.Httpd)
}
