package main

import (
	"flag"
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/files/pdf"
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

var id, path, mongo string

func main() {
	flag.StringVar(&id, "id", "", "Order ID to lookup.")
	flag.StringVar(&path, "path", "invoice.pdf", "Output file path.")
	flag.StringVar(&mongo, "mongo", "localhost/test", "MongoDB to use.")
	flag.Parse()

	o := model.Order{}
	o.Id = store.ParseId(id)

	if !o.Id.Valid() {
		log.Fatal("bad order id")
	}

	s, err := mgo.Dial(mongo)
	if err != nil {
		log.Fatal(err)
	}
	db := s.DB("")

	if err := db.C("orders").FindId(o.Id).One(&o); err != nil {
		log.Fatal(err)
	}
	for i, oi := range o.Items {
		if err := db.FindRef(oi.ProductRef).One(&o.Items[i].Product); err != nil {
			log.Fatal(err)
		}
	}
	if err := db.FindRef(o.CustomerRef).One(&o.Customer); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := pdf.NewInvoice(o).WriteTo(f); err != nil {
		log.Fatal(err)
	}
}
