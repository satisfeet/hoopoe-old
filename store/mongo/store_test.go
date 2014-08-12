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

func TestSuite(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	id    bson.ObjectId
	url   string
	store *Store
	mongo *mgo.Database
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.store = &Store{
		sess,
		sess.DB(""),
	}

	s.mongo = sess.DB("")
}

func (s *Suite) SetUpTest(c *check.C) {
	s.id = bson.NewObjectId()

	err := s.mongo.C("models").Insert(models[0])
	c.Assert(err, check.IsNil)
	err = s.mongo.C("models").Insert(models[1])
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownSuite(c *check.C) {
	s.mongo.Session.Close()
}

func (s *Suite) TearDownTest(c *check.C) {
	_, err := s.mongo.C("models").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestStoreDialAndClose(c *check.C) {
	store := &Store{}

	err := store.Dial(s.url)
	c.Assert(err, check.IsNil)
	err = store.Close()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestStoreFind(c *check.C) {
	m := []model{}

	err := s.store.Find("models", Query{}, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models)
}

func (s *Suite) TestStoreFindId(c *check.C) {
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

func (s *Suite) TestStoreFindOne(c *check.C) {
	q := Query{}
	q.Id(models[0].Id)

	m := model{}

	err := s.store.FindOne("models", q, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}

func (s *Suite) TestStoreInsert(c *check.C) {
	m1 := model{Name: "Baz"}
	m2 := model{}

	err := s.store.Insert("models", &m1)
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").FindId(m1.Id).One(&m2)
	c.Assert(err, check.IsNil)

	c.Check(m1, check.DeepEquals, m2)
}

func (s *Suite) TestStoreUpdate(c *check.C) {
	m := model{}

	models[0].Name += "Foo"

	err := s.store.Update("models", Query{"name": "Foo"}, models[0])
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").FindId(models[0].Id).One(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[0])
}

func (s *Suite) TestStoreUpdateId(c *check.C) {
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

func (s *Suite) TestStoreRemove(c *check.C) {
	m := []model{}

	err := s.store.Remove("models", Query{"name": "Bar"})
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").Find(nil).All(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[:1])
}

func (s *Suite) TestStoreRemoveId(c *check.C) {
	m := []model{}

	err := s.store.RemoveId("models", models[1].Id)
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").Find(nil).All(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, models[:1])
}

func (s *Suite) TestStoreRemoveAll(c *check.C) {
	m := []model{}

	err := s.store.RemoveAll("models", Query{})
	c.Assert(err, check.IsNil)

	err = s.mongo.C("models").Find(nil).All(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, []model{})
}
