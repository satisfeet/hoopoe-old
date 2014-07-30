package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/satisfeet/hoopoe/conf"
	"github.com/satisfeet/hoopoe/httpd"
	"github.com/satisfeet/hoopoe/store"
)

func main() {
	c := conf.NewConf()
	s := store.NewStore()

	if err := c.Flags(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing arguments: %s.\n", err)

		return
	}
	if err := s.Dial(c.Mongo); err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	h := httpd.NewHandler(s)
	h.Auth.Username = c.Username
	h.Auth.Password = c.Password

	if err := http.ListenAndServe(c.Host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}
