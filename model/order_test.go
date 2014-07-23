package model

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestOrder(t *testing.T) {
	check.Suite(&OrderSuite{})
	check.TestingT(t)
}

func TestOrderItem(t *testing.T) {
	check.Suite(&OrderSuite{})
	check.TestingT(t)
}

type OrderSuite struct{}

type OrderItemSuite struct{}

func (s *OrderSuite) TestValidate(c *check.C) {
	i := []OrderItem{
		OrderItem{
			Product: mgo.DBRef{
				Id: bson.NewObjectId(),
			},
			Pricing: Pricing{
				Retail: 299,
			},
			Quantity: 1,
			Variation: Variation{
				Size:  "38",
				Color: "blue",
			},
		},
	}

	c.Check(Order{
		Customer: mgo.DBRef{
			Id: bson.NewObjectId(),
		},
		Items: i,
		Pricing: Pricing{
			Retail: 299,
		},
	}.Validate(), check.IsNil)
}

func (s *OrderItemSuite) TestValidate(c *check.C) {
	p := Pricing{
		Retail: 299,
	}
	v := Variation{
		Color: "black",
		Size:  "38",
	}

	c.Check(OrderItem{
		Product: mgo.DBRef{
			Id: bson.NewObjectId(),
		},
		Quantity:  1,
		Pricing:   p,
		Variation: v,
	}.Validate(), check.IsNil)

	c.Check(OrderItem{
		Quantity:  1,
		Pricing:   p,
		Variation: v,
	}.Validate(), check.ErrorMatches, "product has invalid .*")

	c.Check(OrderItem{
		Product: mgo.DBRef{
			Id: bson.NewObjectId(),
		},
		Quantity:  1,
		Variation: v,
	}.Validate(), check.ErrorMatches, "pricing has invalid .*")

	c.Check(OrderItem{
		Product: mgo.DBRef{
			Id: bson.NewObjectId(),
		},
		Quantity:  0,
		Pricing:   p,
		Variation: v,
	}.Validate(), check.ErrorMatches, "quantity has invalid .*")

	c.Check(OrderItem{
		Product: mgo.DBRef{
			Id: bson.NewObjectId(),
		},
		Quantity: 1,
		Pricing:  p,
	}.Validate(), check.ErrorMatches, "variation has invalid .*")
}
