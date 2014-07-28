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
	if err := conf.Flags(os.Args[1:]); err != nil {
		fmt.Printf("Error parsing arguments: %s.\n", err)

		return
	}
	if err := store.Open(conf.Mongo); err != nil {
		fmt.Printf("Error connecting to database: %s.\n", err)

		return
	}

	h := httpd.Handler(conf.Username, conf.Password)

	if err := http.ListenAndServe(conf.Host, h); err != nil {
		fmt.Printf("Error starting http server: %s.\n", err)
	}
}
