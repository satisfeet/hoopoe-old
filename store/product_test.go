package store

import (
	"io/ioutil"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
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

var ps = &ProductSuite{
	StoreSuite: ss,
}

func TestProduct(t *testing.T) {
	check.Suite(ps)
	check.TestingT(t)
}

type ProductSuite struct {
	*StoreSuite
	store *Product
}

func (s *ProductSuite) SetUpSuite(c *check.C) {
	s.StoreSuite.SetUpSuite(c)

	s.store = NewProduct(s.mongo)
}

func (s *ProductSuite) SetUpTest(c *check.C) {
	err := s.mongo.Insert("products", &products[0])
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) SkipTestProductImage(c *check.C) {
	f, err := s.store.CreateImage(&products[0])
	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)

	err = s.store.FindOne(&products[0])
	c.Assert(err, check.IsNil)

	f, err = s.store.OpenImage(&products[0], products[0].Images[0])
	c.Assert(err, check.IsNil)
	b, err := ioutil.ReadAll(f)
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
	c.Check(b, check.DeepEquals, []byte("Hello"))

	err = s.store.RemoveImage(&products[0], products[0].Images[0])
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TearDownTest(c *check.C) {
	err := s.mongo.RemoveAll("products", mongo.Query{})
	c.Assert(err, check.IsNil)
	err = s.mongo.RemoveAll("products.files", mongo.Query{})
	c.Assert(err, check.IsNil)
	err = s.mongo.RemoveAll("products.chunks", mongo.Query{})
	c.Assert(err, check.IsNil)
}
