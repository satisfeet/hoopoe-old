package main

import (
  "log"
  "fmt"

  "github.com/satisfeet/hoopoe/lib/conf"
)

func main() {
  c, err := conf.New()

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%#v", c)
}
