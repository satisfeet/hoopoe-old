package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

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
		fmt.Print("The host argument is required.\n")

		os.Exit(1)
	}
	if len(conf.Mongo) == 0 {
		fmt.Print("The mongo argument is required.\n")

		os.Exit(1)
	}

	s := NewStore()

	if err := s.Open(conf.Mongo); err != nil {
		fmt.Print("Connection to mongodb failed.\n")

		os.Exit(1)
	}

	http.Handle("/customers", Auth(NewCustomersHandler(s)))
	http.Handle("/", Auth(http.HandlerFunc(NotFound)))

	log.Fatal(http.ListenAndServe(conf.Host, nil))
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
