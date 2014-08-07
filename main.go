package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-handler"
	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
)

func main() {
	c := conf.NewConf()

	if err := c.Flags(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing arguments: %s.\n", err)

		return
	}

	s, err := mgo.Dial(c.Mongo)

	if err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	m := http.NewServeMux()
	m.Handle("/", httpd.NewCustomerHandler(s.DB("")))

	h := &handler.Logger{
		Handler: &handler.Auth{
			Username: c.Username,
			Password: c.Password,
			Handler:  m,
		},
	}

	if err := http.ListenAndServe(c.Host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}
