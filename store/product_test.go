package store

import (
	"io/ioutil"

	"gopkg.in/check.v1"

	"github.com/satisfeet/hoopoe/model"
)

var products = []model.Product{
	model.Product{
		Name: "Summer socks",
		Pricing: model.Pricing{
			Retail: 599,
		},
		Variations: []model.Variation{
			model.Variation{
				Size:  "42-44",
				Color: "black",
			},
		},
	},
}

func (s *Suite) SkipTestProductImage(c *check.C) {
	f, err := s.product.CreateImage(products[0].Id)
	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)

	err = s.product.FindId(products[0].Id, &products[0])
	c.Assert(err, check.IsNil)

	f, err = s.product.OpenImage(products[0].Id, products[0].Images[0])
	c.Assert(err, check.IsNil)
	b, err := ioutil.ReadAll(f)
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
	c.Check(b, check.DeepEquals, []byte("Hello"))

	err = s.product.RemoveImage(products[0].Id, products[0].Images[0])
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestProductRemoveId(c *check.C) {
	err := s.customer.RemoveId(customers[0].Id)
	c.Assert(err, check.IsNil)
}
