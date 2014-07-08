package main

import (
	"log"
	"net/http"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

func main() {
	c := conf.NewConf()
	s := store.NewStore()
	h := httpd.NewHttpd(s)

	if err := c.ParseFlags(); err != nil {
		log.Fatal(err)
	}

	if err := s.Open(c["mongo"]); err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(c["addr"], h)

	log.Printf("Hoopoe listening on %s", c["addr"])
}
