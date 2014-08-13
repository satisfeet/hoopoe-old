package store

import (
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/model"
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

func (s *Suite) TestOrderFindCustomer(c *check.C) {
	err := s.order.FindCustomer(&order)
	c.Assert(err, check.IsNil)

	c.Check(order.Customer, check.DeepEquals, customers[0])
}

func (s *Suite) TestOrderFindProducts(c *check.C) {
	err := s.order.FindProducts(&order)
	c.Assert(err, check.IsNil)

	c.Check(order.Items[0].Product, check.DeepEquals, products[0])
}
