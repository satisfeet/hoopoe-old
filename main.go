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

	if err := store.Open(conf.Mongo); err != nil {
		fmt.Print("Connection to mongodb failed.\n")

		os.Exit(1)
	}

	http.Handle("/customers", httpd.Auth(httpd.NewCustomers()))
	http.Handle("/", httpd.Auth(NotFound()))

	log.Fatal(http.ListenAndServe(conf.Host, nil))
}

func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})
}
