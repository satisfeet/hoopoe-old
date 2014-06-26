package httpd

import (
  "net/http"

  "github.com/satisfeet/hoopoe/lib/conf"
)

func Listen(c *conf.HttpdConfig) {
  http.HandleFunc("/", handle)
  http.ListenAndServe(c.Port, nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(200)
  w.Write([]byte("Hello World"))
}
