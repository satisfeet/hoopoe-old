package store

import (
	"io/ioutil"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var products = []model.Product{
	model.Product{
		Id:     bson.NewObjectId(),
		Name:   "Summer socks",
		Images: []bson.ObjectId{},
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
	f, err := s.product.CreateImage(&products[0])
	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)

	err = s.product.FindOne(&products[0])
	c.Assert(err, check.IsNil)

	f, err = s.product.OpenImage(&products[0], products[0].Images[0])
	c.Assert(err, check.IsNil)
	b, err := ioutil.ReadAll(f)
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
	c.Check(b, check.DeepEquals, []byte("Hello"))

	err = s.product.RemoveImage(&products[0], products[0].Images[0])
	c.Assert(err, check.IsNil)
}
