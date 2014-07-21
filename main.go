package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

var (
	conf struct {
		Host  string
		Mongo string
	}
)

func main() {
	s := &store.Session{}

	flag.StringVar(&conf.Host, "host", "", "The http host to use.")
	flag.StringVar(&conf.Mongo, "mongo", "", "The mongodb url to use.")
	flag.Parse()

	if len(conf.Host) == 0 {
		fmt.Print("The host argument is required.\n")

		os.Exit(1)
	}
	if len(conf.Mongo) == 0 {
		fmt.Print("The mongo argument is required.\n")

		os.Exit(1)
	}

	if err := s.Open(conf.Mongo); err != nil {
		fmt.Print("Connection to mongodb failed.\n")

		os.Exit(1)
	}

	http.Handle("/customers", httpd.Auth(&httpd.Customers{
		Store: &store.Store{
			Name:    "customers",
			Session: s,
		},
	}))
	http.Handle("/", httpd.Auth(httpd.NotFound()))

	log.Fatal(http.ListenAndServe(conf.Host, nil))
}
