package main

import (
    "log"

    "github.com/satisfeet/hoopoe/lib/conf"
    "github.com/satisfeet/hoopoe/lib/store"
    "github.com/satisfeet/hoopoe/lib/httpd"
)

const (
  DEFAULT = "/etc/default.json"
  DEVELOPMENT = "/etc/development.json"
)

func main() {
    c := conf.New()

    if err := c.LoadJSON(DEFAULT); err != nil {
        log.Fatal(err)
    }
    if err := c.LoadJSON(DEVELOPMENT); err != nil {
        log.Fatal(err)
    }

    s := store.New()

    if err := s.Open(c); err != nil {
        log.Fatal(err)
    }

    httpd.Listen(c)
}
