package mongo

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type model struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string        `store:"index"`
}

var models = []model{
	model{bson.NewObjectId(), "Foo"},
	model{bson.NewObjectId(), "Bar"},
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
	store *Store
	mongo *mgo.Database
}

func (s *StoreSuite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.store = &Store{
		sess,
		sess.DB(""),
	}

	s.mongo = sess.DB("")
}

func (s *StoreSuite) SetUpTest(c *check.C) {
	err := s.mongo.C("models").Insert(models[0])
	c.Assert(err, check.IsNil)
	err = s.mongo.C("models").Insert(models[1])
	c.Assert(err, check.IsNil)
}

func (s *StoreSuite) TestDialAndClose(c *check.C) {
	store := &Store{}

	err := store.Dial(s.url)
	c.Assert(err, check.IsNil)
	err = store.Close()
	c.Assert(err, check.IsNil)
}

func (s *StoreSuite) TestFind(c *check.C) {
	m := []model{}

	err := s.store.Find("models", Query{}, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models)
}

func (s *StoreSuite) TestFindId(c *check.C) {
	m1 := model{}
	m2 := model{}
	m3 := model{}

	err := s.store.FindId("models", "1234", &m1)
	c.Assert(err, check.Equals, ErrBadId)
	err = s.store.FindId("models", models[0].Id, &m2)
	c.Assert(err, check.IsNil)
	err = s.store.FindId("models", models[0].Id.Hex(), &m3)
	c.Assert(err, check.IsNil)

	c.Check(m1, check.DeepEquals, model{})
	c.Check(m2, check.DeepEquals, models[0])
	c.Check(m3, check.DeepEquals, models[0])
}

func (s *StoreSuite) TestFindOne(c *check.C) {
	q := Query{}
	q.Id(models[0].Id)

	m := model{}

	err := s.store.FindOne("models", q, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}

func (s *StoreSuite) TestInsert(c *check.C) {
	m1 := model{Name: "Baz"}
	m2 := model{}

	err := s.store.Insert("models", &m1)
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").FindId(m1.Id).One(&m2)
	c.Assert(err, check.IsNil)

	c.Check(m1, check.DeepEquals, m2)
}

func (s *StoreSuite) TestUpdate(c *check.C) {
	m := model{}

	models[0].Name = "FooFoo"

	err := s.store.Update("models", Query{"_id": models[0].Id}, models[0])
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").FindId(models[0].Id).One(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}

func (s *StoreSuite) TestUpdateId(c *check.C) {
	m := model{}

	models[0].Name += "Foo"

	err := s.store.UpdateId("models", models[0].Id.Hex(), models[0])
	c.Assert(err, check.IsNil)
	err = s.store.UpdateId("models", models[0].Id.Hex()+"0", models[0])
	c.Assert(err, check.Equals, ErrBadId)
	err = s.store.UpdateId("models", "12345", models[0])
	c.Assert(err, check.Equals, ErrBadId)

	err = s.mongo.C("models").FindId(models[0].Id).One(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}

func (s *StoreSuite) TestRemove(c *check.C) {
	m := []model{}

	err := s.store.Remove("models", Query{"name": "Bar"})
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").Find(nil).All(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[:1])
}

func (s *StoreSuite) TestRemoveId(c *check.C) {
	m := []model{}

	err := s.store.RemoveId("models", models[1].Id)
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").Find(nil).All(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[:1])
}

func (s *StoreSuite) TestRemoveAll(c *check.C) {
	m := []model{}

	err := s.store.RemoveAll("models", Query{})
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").Find(nil).All(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, []model{})
}

func (s *StoreSuite) TearDownTest(c *check.C) {
	_, err := s.mongo.C("models").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *StoreSuite) TearDownStoreSuite(c *check.C) {
	s.mongo.Session.Close()
}
