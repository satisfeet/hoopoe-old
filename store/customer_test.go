package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/mongo"
)

func TestCustomer(t *testing.T) {
	check.Suite(&CustomerSuite{})
	check.Suite(&CustomersSuite{
		Url: "localhost/test",
	})
	check.TestingT(t)
}

type CustomerSuite struct{}

func (s *CustomerSuite) TestValidate(c *check.C) {
	a := Address{
		City: "Berlin",
	}

	c.Check(Customer{
		Name:    "Bodo Kaiser",
		Email:   "i@bodokaiser.io",
		Address: a,
	}.Validate(), check.IsNil)
	c.Check(Customer{
		Name:    "Bodo Kaiser",
		Email:   "i@bodokaiser.io",
		Address: a,
		Company: "satisfeet",
	}.Validate(), check.IsNil)

	c.Check(Customer{
		Email:   "foo@bar.org",
		Address: a,
	}.Validate(), check.ErrorMatches, "Name .*")
	c.Check(Customer{
		Name:    "Bodo Kaiser",
		Address: a,
	}.Validate(), check.ErrorMatches, "Email .*")
	c.Check(Customer{
		Name:  "Bodo Kaiser",
		Email: "foo@bar.org",
	}.Validate(), check.ErrorMatches, "Address.*")
}

type CustomersSuite struct {
	Url       string
	Customer  *Customer
	Customers *Customers
}

func (s *CustomersSuite) SetUpSuite(c *check.C) {
	s.Customers = &Customers{
		Mongo: &mongo.Store{},
	}
	c.Assert(s.Customers.Mongo.Open(s.Url), check.IsNil)
}

func (s *CustomersSuite) SetUpTest(c *check.C) {
	s.Customer = &Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
		Address: Address{
			City: "Berlin",
		},
	}
	c.Assert(s.Customers.Mongo.Insert(CustomersName, s.Customer), check.IsNil)
}

func (s *CustomersSuite) TestOne(c *check.C) {
	m := &Customer{}

	c.Check(s.Customers.One(s.Customer.Id.Hex(), m), check.IsNil)
	c.Check(m, check.DeepEquals, s.Customer)
}

func (s *CustomersSuite) TestAll(c *check.C) {
	m := []Customer{}

	c.Check(s.Customers.All(&m), check.IsNil)
	c.Check(m, check.DeepEquals, []Customer{*s.Customer})
}

func (s *CustomersSuite) TestSearch(c *check.C) {
	m1 := []Customer{}
	m2 := []Customer{}

	c.Check(s.Customers.Search("aise", &m1), check.IsNil)
	c.Check(s.Customers.Search("123shfajd", &m2), check.IsNil)
	c.Check(m1, check.DeepEquals, []Customer{*s.Customer})
	c.Check(m2, check.DeepEquals, []Customer{})
}

func (s *CustomersSuite) TestInsert(c *check.C) {
	m1 := &Customer{
		Name:  "Will Smith",
		Email: "will@bel-air.us",
		Address: Address{
			City: "Los Angeles",
		},
	}
	m2 := &Customer{}

	c.Check(s.Customers.Insert(m1), check.IsNil)

	q := mongo.Query{"_id": m1.Id}
	c.Check(s.Customers.Mongo.FindOne(CustomersName, q, m2), check.IsNil)
	c.Check(m2, check.DeepEquals, m1)
}

func (s *CustomersSuite) TestUpdate(c *check.C) {
	s.Customer.Name += " Junior"
	c.Check(s.Customers.Update(s.Customer), check.IsNil)

	m := &Customer{}
	q := mongo.Query{"_id": s.Customer.Id}
	c.Check(s.Customers.Mongo.FindOne(CustomersName, q, m), check.IsNil)
	c.Check(m, check.DeepEquals, s.Customer)
}

func (s *CustomersSuite) TestRemove(c *check.C) {
	c.Check(s.Customers.Remove(s.Customer), check.IsNil)

	q := mongo.Query{"_id": s.Customer.Id}
	c.Check(s.Customers.Mongo.FindOne(CustomersName, q, nil), check.Equals, mgo.ErrNotFound)
}

func (s *CustomersSuite) TearDownTest(c *check.C) {
	c.Assert(s.Customers.Mongo.Drop(CustomersName), check.IsNil)
}

func (s *CustomersSuite) TearDownSuite(c *check.C) {
	c.Assert(s.Customers.Mongo.Close(), check.IsNil)
}
