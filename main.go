package main

import (
	"flag"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/httpd"
)

var host, mongo, username, password string

func main() {
	flag.StringVar(&username, "username", "bodokaiser", "")
	flag.StringVar(&password, "password", "secret", "")
	flag.StringVar(&mongo, "mongo", "localhost/satisfeet", "")
	flag.StringVar(&host, "host", ":3000", "")
	flag.Parse()

	s, err := mgo.Dial(mongo)

	if err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	m := http.NewServeMux()
	m.Handle("/", httpd.NewCustomerHandler(s.DB("")))

	h := &handler.Logger{
		Handler: &handler.Auth{
			Username: username,
			Password: password,
			Handler:  m,
		},
	}

	if err := http.ListenAndServe(host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}
