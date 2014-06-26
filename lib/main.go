package main

import (
  "log"

  "github.com/satisfeet/hoopoe/lib/conf"
  "github.com/satisfeet/hoopoe/lib/httpd"
)

func main() {
  c, err := conf.New()

  if err != nil {
    log.Fatal(err)
  }

  httpd.Listen(&c.Httpd)
}
