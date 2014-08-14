package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/mongo"
)

type Model struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `store:"index"`
}

var Models = []Model{
	Model{bson.NewObjectId(), "Foo"},
	Model{bson.NewObjectId(), "Bar"},
}

var ss = &StoreSuite{
	url: "localhost/test",
}

func TestStore(t *testing.T) {
	check.Suite(ss)
	check.TestingT(t)
}

type StoreSuite struct {
	url   string
	mongo *mongo.Store
	store *store
}

func (s *StoreSuite) SetUpSuite(c *check.C) {
	s.mongo = &mongo.Store{}

	err := s.mongo.Dial(s.url)
	c.Assert(err, check.IsNil)
}

func (s *StoreSuite) SetUpTest(c *check.C) {
	s.store = &store{s.mongo}

	err := s.mongo.Insert("models", &Models[0])
	c.Assert(err, check.IsNil)
	err = s.mongo.Insert("models", &Models[1])
	c.Assert(err, check.IsNil)
}

func (s *StoreSuite) TestFind(c *check.C) {
	m := []Model{}

	if s.store == nil {
		return
	}

	err := s.store.Find(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, Models)
}

func (s *StoreSuite) TestFindOne(c *check.C) {
	m := Model{Id: Models[0].Id}

	if s.store == nil {
		return
	}

	err := s.store.FindOne(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, Models[0])
}

func (s *StoreSuite) TestInsert(c *check.C) {
	m1 := Model{Name: "Bodo"}
	m2 := Model{}

	if s.store == nil {
		return
	}

	err := s.store.Insert(&m1)
	c.Assert(err, check.IsNil)

	err = s.mongo.FindId("models", m1.Id, &m2)
	c.Assert(err, check.IsNil)

	c.Check(m1, check.DeepEquals, m2)
}

func (s *StoreSuite) TestUpdate(c *check.C) {
	m := Model{}

	Models[0].Name += "Foo"

	if s.store == nil {
		return
	}

	err := s.store.Update(Models[0])
	c.Assert(err, check.IsNil)

	err = s.mongo.FindId("models", Models[0].Id, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, Models[0])
}

func (s *StoreSuite) TearDownTest(c *check.C) {
	if s.store == nil {
		return
	}

	err := s.mongo.RemoveAll("models", mongo.Query{})
	c.Assert(err, check.IsNil)
}

func (s *StoreSuite) TearDownSuite(c *check.C) {
	err := s.mongo.Close()
	c.Assert(err, check.IsNil)
}
