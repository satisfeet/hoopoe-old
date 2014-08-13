package main

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
)

var order = model.Order{
	Id: bson.NewObjectId(),
	Items: []model.OrderItem{
		model.OrderItem{
			ProductRef: &mgo.DBRef{
				Id:         product.Id,
				Collection: "products",
			},
			Pricing:   pricing,
			Variation: product.Variations[0],
			Quantity:  1,
		},
	},
	Pricing: pricing,
	CustomerRef: &mgo.DBRef{
		Id:         customer.Id,
		Collection: "customers",
	},
}

var pricing = model.Pricing{
	Retail: 599,
}

var product = model.Product{
	Id:      bson.NewObjectId(),
	Name:    "Summer socks",
	Pricing: pricing,
	Variations: []model.Variation{
		model.Variation{
			Size:  "42-44",
			Color: "black",
		},
	},
}

var customer = model.Customer{
	Id:    bson.NewObjectId(),
	Name:  "Haci Erdal",
	Email: "haci@hotmail.de",
	Address: model.Address{
		Street:  "Checkpoint Charlie 2",
		City:    "Berlin",
		ZipCode: 11001,
	},
}

func TestMain(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	url   string
	mongo *mongo.Store
}

func (s *Suite) SetUpSuite(c *check.C) {
	s.mongo = &mongo.Store{}

	err := s.mongo.Dial(s.url)
	c.Assert(err, check.IsNil)

	err = s.mongo.Insert("orders", &order)
	c.Assert(err, check.IsNil)
	err = s.mongo.Insert("products", &product)
	c.Assert(err, check.IsNil)
	err = s.mongo.Insert("customers", &customer)
	c.Assert(err, check.IsNil)

	os.Args = append(os.Args, []string{"--order", order.Id.Hex()}...)
}

func (s *Suite) TearDownSuite(c *check.C) {
	err := s.mongo.RemoveAll("orders", mongo.Query{})
	c.Assert(err, check.IsNil)
	err = s.mongo.RemoveAll("products", mongo.Query{})
	c.Assert(err, check.IsNil)
	err = s.mongo.RemoveAll("customers", mongo.Query{})
	c.Assert(err, check.IsNil)

	err = s.mongo.Close()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestMain(c *check.C) {
	main()

	f, err := os.Open("invoice.pdf")
	c.Assert(err, check.IsNil)
	f.Close()

	cmd := exec.Command("open", "invoice.pdf")
	err = cmd.Start()
	c.Assert(err, check.IsNil)
	err = cmd.Wait()
	c.Assert(err, check.IsNil)

	time.Sleep(2 * time.Second)

	err = os.Remove("invoice.pdf")
	c.Assert(err, check.IsNil)
}
