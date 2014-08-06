package store

import (
	"testing"

	"github.com/satisfeet/hoopoe/store/mongo"
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestStore(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	url        string
	model      model
	store      *Store
	database   *mgo.Database
	collection *mgo.Collection
}

type model struct {
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `store:"index"`
	Email string        `store:"unique"`
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	m := &mongo.Store{}
	err = m.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.store = &Store{m}
	s.database = sess.DB("")
	s.collection = sess.DB("").C("models")
}

func (s *Suite) SetUpTest(c *check.C) {
	s.model = model{
		Id:    bson.NewObjectId(),
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
	}

	err := s.collection.Insert(&s.model)
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestDialAndClose(c *check.C) {
	store := NewStore()

	c.Assert(store.Dial(s.url), check.IsNil)
	c.Assert(store.Dial(s.url), check.NotNil)
	c.Assert(store.Close(), check.IsNil)
	c.Assert(store.Close(), check.NotNil)
}

func (s *Suite) TestInsert(c *check.C) {
	m1 := model{
		Name:  "Jason Statham",
		Email: "j@stats.uk",
	}
	m2 := model{}

	err := s.store.Insert(&m1)
	c.Assert(err, check.IsNil)

	err = s.collection.FindId(m1.Id).One(&m2)
	c.Assert(err, check.IsNil)
	c.Check(m1, check.DeepEquals, m2)
}

func (s *Suite) TestUpdate(c *check.C) {
	m := model{}
	s.model.Name += " I"

	err := s.store.Update(s.model)
	c.Assert(err, check.IsNil)

	err = s.collection.FindId(s.model.Id).One(&m)
	c.Assert(err, check.IsNil)
	c.Check(m, check.DeepEquals, s.model)
}

func (s *Suite) TestRemove(c *check.C) {
	err := s.store.Remove(s.model)
	c.Assert(err, check.IsNil)

	err = s.collection.FindId(s.model.Id).One(nil)
	c.Check(err, check.Equals, mgo.ErrNotFound)
}

func (s *Suite) TearDownTest(c *check.C) {
	_, err := s.collection.RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownSuite(c *check.C) {
	s.database.Session.Close()
}
