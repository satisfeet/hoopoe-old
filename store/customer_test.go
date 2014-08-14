package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var customer = model.Customer{
	Id:    bson.NewObjectId(),
	Name:  "Bodo Kaiser",
	Email: "i@bodokaiser.io",
	Address: model.Address{
		City: "Berlin",
	},
}

func TestCustomer(t *testing.T) {
	check.Suite(&CustomerSuite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type CustomerSuite struct {
	url      string
	store    *Customer
	session  *mgo.Session
	database *mgo.Database
}

func (s *CustomerSuite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.session = sess
	s.database = sess.DB("")

	s.store = NewCustomer(sess)
}

func (s *CustomerSuite) SetUpTest(c *check.C) {
	err := s.database.C("customers").Insert(customer)
	c.Assert(err, check.IsNil)
}

func (s *CustomerSuite) TestFind(c *check.C) {
	m := []model.Customer{}

	err := s.store.Find(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.HasLen, 1)
	c.Check(m[0], check.DeepEquals, customer)
}

func (s *CustomerSuite) TestFindOne(c *check.C) {
	m1 := model.Customer{Id: customer.Id}
	m2 := model.Customer{Id: bson.ObjectId("1234")}

	err1 := s.store.FindOne(&m1)
	err2 := s.store.FindOne(&m2)

	c.Assert(err1, check.IsNil)
	c.Assert(err2, check.Equals, ErrBadId)

	c.Check(m1, check.DeepEquals, customer)
}

func (s *CustomerSuite) TestSearch(c *check.C) {
	m1 := []model.Customer{}
	m2 := []model.Customer{}

	err1 := s.store.Search("ais", &m1)
	err2 := s.store.Search("foo", &m2)

	c.Assert(err1, check.IsNil)
	c.Assert(err2, check.IsNil)

	c.Check(m1, check.HasLen, 1)
	c.Check(m1[0], check.DeepEquals, customer)
	c.Check(m2, check.HasLen, 0)
}

func (s *CustomerSuite) TestInsert(c *check.C) {
	m := model.Customer{
		Name:  "Denzel Washington",
		Email: "denzel@gmail.com",
		Address: model.Address{
			City: "Los Angelos",
		},
	}

	err := s.store.Insert(&m)
	c.Assert(err, check.IsNil)

	c.Check(m.Id.Valid(), check.Equals, true)
}

func (s *CustomerSuite) TestUpdate(c *check.C) {
	customer.Name += " Jr"

	err := s.store.Update(&customer)
	c.Assert(err, check.IsNil)
}

func (s *CustomerSuite) TestRemove(c *check.C) {
	err := s.store.Remove(customer)
	c.Assert(err, check.IsNil)
}

func (s *CustomerSuite) TearDownTest(c *check.C) {
	_, err := s.database.C("customers").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *CustomerSuite) TearDownSuite(c *check.C) {
	s.session.Close()
}
