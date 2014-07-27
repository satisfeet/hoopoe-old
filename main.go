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
	c := &conf.Conf{}

	if err := c.Flags(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing arguments: %s.\n", err)

		return
	}
	if err := store.Open(c.Mongo); err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	h := httpd.Handler(c.Username, c.Password)

	if err := http.ListenAndServe(c.Host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}
