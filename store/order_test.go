package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
)

var order = model.Order{
	Items: []model.OrderItem{
		model.OrderItem{
			ProductRef: &mgo.DBRef{
				Id:         products[0].Id,
				Collection: "products",
			},
			Pricing: model.Pricing{
				Retail: 599,
			},
			Variation: products[0].Variations[0],
			Quantity:  1,
		},
	},
	Pricing: model.Pricing{
		Retail: 599,
	},
	CustomerRef: &mgo.DBRef{
		Id:         customers[0].Id,
		Collection: "customers",
	},
}

var os = &OrderSuite{
	StoreSuite:    ss,
	ProductSuite:  ps,
	CustomerSuite: cs,
}

func TestOrder(t *testing.T) {
	check.Suite(os)
	check.TestingT(t)
}

type OrderSuite struct {
	*StoreSuite
	*ProductSuite
	*CustomerSuite
	store *Order
}

func (s *OrderSuite) SetUpSuite(c *check.C) {
	s.StoreSuite.SetUpSuite(c)
	s.ProductSuite.SetUpSuite(c)
	s.CustomerSuite.SetUpSuite(c)

	s.store = NewOrder(s.mongo)
}

func (s *OrderSuite) SetUpTest(c *check.C) {
	s.ProductSuite.SetUpTest(c)
	s.CustomerSuite.SetUpTest(c)

	err := s.mongo.Insert("orders", &order)
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TestFindCustomer(c *check.C) {
	err := s.store.FindCustomer(&order)
	c.Assert(err, check.IsNil)

	c.Check(order.Customer, check.DeepEquals, customers[0])
}

func (s *OrderSuite) TestFindProducts(c *check.C) {
	err := s.store.FindProducts(&order)
	c.Assert(err, check.IsNil)

	c.Check(order.Items[0].Product, check.DeepEquals, products[0])
}

func (s *OrderSuite) TearDownTest(c *check.C) {
	s.ProductSuite.TearDownTest(c)
	s.CustomerSuite.TearDownTest(c)

	err := s.mongo.RemoveAll("orders", mongo.Query{})
	c.Assert(err, check.IsNil)
}
