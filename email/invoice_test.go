package email

import (
	"bytes"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

func TestPDF(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

var order = model.Order{
	Id: bson.NewObjectId(),
	Items: []model.OrderItem{
		model.OrderItem{
			Product: model.Product{
				Id:   bson.NewObjectId(),
				Name: "Summer socks",
				Pricing: model.Pricing{
					Retail: 599,
				},
			},
			Pricing: model.Pricing{
				Retail: 599,
			},
			Variation: model.Variation{
				Size:  "42-44",
				Color: "black",
			},
			Quantity: 1,
		},
	},
	Pricing: model.Pricing{
		Retail: 599,
	},
	Customer: model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Haci Erdal",
		Email: "haci@hotmail.de",
		Address: model.Address{
			Street:  "Checkpoint Charlie 2",
			City:    "Berlin",
			ZipCode: 11001,
		},
	},
}

type Suite struct{}

func (s *Suite) TestInvoice(c *check.C) {
	b := &bytes.Buffer{}

	_, err := NewInvoice(order).WriteTo(b)
	c.Assert(err, check.IsNil)
}
