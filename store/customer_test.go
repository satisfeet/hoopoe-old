package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
)

var customers = []model.Customer{
	model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
		Address: model.Address{
			City: "Berlin",
		},
	},
	model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Denzel Washington",
		Email: "denzel@example.com",
		Address: model.Address{
			City: "Hollywood",
		},
	},
}

var cs = &CustomerSuite{
	StoreSuite: ss,
}

func TestCustomer(t *testing.T) {
	check.Suite(cs)
	check.TestingT(t)
}

type CustomerSuite struct {
	*StoreSuite
	store *Customer
}

func (s *CustomerSuite) SetUpSuite(c *check.C) {
	s.StoreSuite.SetUpSuite(c)

	s.store = NewCustomer(s.mongo)
}

func (s *CustomerSuite) SetUpTest(c *check.C) {
	err := s.mongo.Insert("customers", &customers[0])
	c.Assert(err, check.IsNil)
	err = s.mongo.Insert("customers", &customers[1])
	c.Assert(err, check.IsNil)
}

func (s *CustomerSuite) TestSearch(c *check.C) {
	m := []model.Customer{}

	err := s.store.Search("denzel", &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, customers[1:])
}

func (s *CustomerSuite) TearDownTest(c *check.C) {
	err := s.mongo.RemoveAll("customers", mongo.Query{})
	c.Assert(err, check.IsNil)
}
