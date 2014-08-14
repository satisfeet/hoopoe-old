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

func main() {
	var orderId, output, mongodb string

	flag.StringVar(&orderId, "order", "", "Order ID to lookup.")
	flag.StringVar(&output, "output", "invoice.pdf", "Output file path.")
	flag.StringVar(&mongodb, "mongo", "localhost/test", "MongoDB to use.")
	flag.Parse()

	o := model.Order{}
	o.Id = store.IdFromString(orderId)

	m, err := mgo.Dial(mongodb)

	if err != nil {
		log.Fatal(err)
	}

	defer m.Close()

	s := store.NewOrder(m)

	if err := s.FindOne(&o); err != nil {
		log.Fatal(err)
	}
	if err := s.FindCustomer(&o); err != nil {
		log.Fatal(err)
	}
	if err := s.FindProducts(&o); err != nil {
		log.Fatal(err)
	}

	if err := writeInvoiceToFile(o, output); err != nil {
		log.Fatal(err)
	}
}

func writeInvoiceToFile(o model.Order, path string) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = pdf.NewInvoice(o).WriteTo(f)

	return err
}
