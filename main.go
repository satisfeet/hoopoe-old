package main

import (
	"log"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

func main() {
	c := conf.New()
	s := store.New()
	h := httpd.New(s)

	if err := s.Open(c.Store); err != nil {
		log.Fatal(err)
	}

	h.Listen(c.Httpd)
}
