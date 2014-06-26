package main

import (
    "log"

    "github.com/satisfeet/hoopoe/lib/conf"
    "github.com/satisfeet/hoopoe/lib/store"
    "github.com/satisfeet/hoopoe/lib/httpd"
)


func main() {
    if err := conf.Init(); err != nil {
        log.Fatal(err)
    }

    if err := store.Init(); err != nil {
        log.Fatal(err)
    }

    if err := httpd.Init(); err != nil {
        log.Fatal(err)
    }
}
