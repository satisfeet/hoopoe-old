package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/mongo"
)

type model struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `store:"index"`
}

var models = []model{
	model{bson.NewObjectId(), "Foo"},
	model{bson.NewObjectId(), "Bar"},
}

func TestSuite(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	url      string
	mongo    *mongo.Store
	store    *store
	customer *CustomerStore
	product  *ProductStore
}

func (s *Suite) SetUpSuite(c *check.C) {
	s.mongo = &mongo.Store{}

	err := s.mongo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.store = &store{s.mongo}
	s.customer = NewCustomerStore(s.mongo)
	s.product = NewProductStore(s.mongo)
}

func (s *Suite) SetUpTest(c *check.C) {
	err := s.mongo.Insert("models", &models[0])
	c.Assert(err, check.IsNil)
	err = s.mongo.Insert("models", &models[1])
	c.Assert(err, check.IsNil)

	err = s.mongo.Insert("products", &products[0])
	c.Assert(err, check.IsNil)

	err = s.mongo.Insert("customers", &customers[0])
	c.Assert(err, check.IsNil)
	err = s.mongo.Insert("customers", &customers[1])
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownTest(c *check.C) {
	err := s.mongo.RemoveAll("models", mongo.Query{})
	c.Assert(err, check.IsNil)

	err = s.mongo.RemoveAll("products", mongo.Query{})
	c.Assert(err, check.IsNil)
	err = s.mongo.RemoveAll("products.files", mongo.Query{})
	c.Assert(err, check.IsNil)
	err = s.mongo.RemoveAll("products.chunks", mongo.Query{})
	c.Assert(err, check.IsNil)

	err = s.mongo.RemoveAll("customers", mongo.Query{})
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownSuite(c *check.C) {
	err := s.mongo.Close()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestStoreFind(c *check.C) {
	m := []model{}

	err := s.store.Find(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models)
}

func (s *Suite) TestStoreFindId(c *check.C) {
	m := model{}

	err := s.store.FindId(models[0].Id, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}

func (s *Suite) TestStoreInsert(c *check.C) {
	m1 := model{Name: "Bodo"}
	m2 := model{}

	err := s.store.Insert(&m1)
	c.Assert(err, check.IsNil)

	err = s.mongo.FindId("models", m1.Id, &m2)
	c.Assert(err, check.IsNil)

	c.Check(m1, check.DeepEquals, m2)
}

func (s *Suite) TestStoreUpdate(c *check.C) {
	m := model{}

	models[0].Name += "Foo"

	err := s.store.Update(models[0])
	c.Assert(err, check.IsNil)

	err = s.mongo.FindId("models", models[0].Id, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}
