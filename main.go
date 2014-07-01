package main

import (
	"log"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

func main() {
	c := conf.New()

	if err := store.Init(c.Store); err != nil {
		log.Fatal(err)
	}

	httpd.Init(c.Httpd)
}
