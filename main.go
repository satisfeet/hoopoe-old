package main

import (
	"log"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
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
