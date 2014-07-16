package main

import (
	"errors"
	"flag"
	"log"
	"net/http"

	. "github.com/satisfeet/hoopoe/net/http"
	. "github.com/satisfeet/hoopoe/store"
)

var (
	conf struct {
		Host  string
		Mongo string
	}
)

func main() {
	flag.StringVar(&conf.Host, "host", "", "The http host to use.")
	flag.StringVar(&conf.Mongo, "mongo", "", "The mongodb url to use.")
	flag.Parse()

	if len(conf.Host) == 0 {
		log.Fatal(errors.New("The host argument is required."))
	}
	if len(conf.Mongo) == 0 {
		log.Fatal(errors.New("The mongo argument is required."))
	}

	s := NewStore()

	if err := s.Open(conf.Mongo); err != nil {
		log.Fatal(err)
	}

	http.Handle("/customers", NewCustomersHandler(s))
	http.ListenAndServe(conf.Host, nil)
}
