package main

import (
	"flag"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
)

var host, auth, mongo string

func main() {
	flag.StringVar(&host, "host", ":3000", "")
	flag.StringVar(&auth, "auth", "bodokaiser:secret", "")
	flag.StringVar(&mongo, "mongo", "localhost/satisfeet", "")
	flag.Parse()

	s, err := mgo.Dial(mongo)
	if err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	m := http.NewServeMux()
	m.Handle("/", httpd.NewCustomerHandler(s.DB("")))

	h := handler.Logger(m)
	h = handler.Auth(auth, h)

	if err := http.ListenAndServe(host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}
