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
)

var order = model.Order{
	Id: bson.NewObjectId(),
	Items: []model.OrderItem{
		model.OrderItem{
			ProductId: product.Id,
			Pricing:   pricing,
			Variation: product.Variations[0],
			Quantity:  1,
		},
	},
	Pricing:    pricing,
	CustomerId: customer.Id,
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
	url      string
	session  *mgo.Session
	database *mgo.Database
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.session = sess
	s.database = sess.DB("")

	err = s.database.C("orders").Insert(&order)
	c.Assert(err, check.IsNil)
	err = s.database.C("products").Insert(&product)
	c.Assert(err, check.IsNil)
	err = s.database.C("customers").Insert(&customer)
	c.Assert(err, check.IsNil)

	os.Args = append(os.Args, []string{"--order", order.Id.Hex()}...)
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

func (s *Suite) TearDownSuite(c *check.C) {
	_, err := s.database.C("orders").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("products").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("customers").RemoveAll(nil)
	c.Assert(err, check.IsNil)

	s.session.Close()
}
