package httpd

import (
  "log"
  "net/http"

  "github.com/satisfeet/hoopoe/lib/conf"
)

func Listen(c *conf.HttpdConfig) {
  m := NewMux()

  m.Use(logger)
  m.Get("/", handle)

  http.ListenAndServe(c.Port, m)
}

func logger(c *Context) {
  log.Printf("Request: %s %s", c.Method(), c.Path())

  c.Next()
}

func handle(c *Context) {
  c.Respond("Hello World", 200)
}
