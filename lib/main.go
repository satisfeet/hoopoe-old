package main

import (
    "log"

    "github.com/satisfeet/hoopoe/lib/app"
    "github.com/satisfeet/hoopoe/lib/httpd"
)

const (
  DEFAULT = "/etc/default.json"
  DEVELOPMENT = "/etc/development.json"
)

func main() {
    a := app.New()

    if err := a.Configure(DEFAULT); err != nil {
        log.Fatal(err)
    }
    if err := a.Configure(DEVELOPMENT); err != nil {
        log.Fatal(err)
    }

    httpd.Listen(a)
}
