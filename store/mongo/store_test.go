package mongo

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestStore(t *testing.T) {
	check.Suite(&StoreSuite{
		Url: "localhost/test",
	})
	check.TestingT(t)
}

type Model struct {
	Id   bson.ObjectId `bson:"_id"`
	Age  int
	Name string
}

type StoreSuite struct {
	Url   string
	Store *Store
	Model *Model
	Mongo *mgo.Database
}

func (s *StoreSuite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.Url)
	c.Assert(err, check.IsNil)

	s.Mongo = sess.DB("")
	s.Store = &Store{sess}
}

func (s *StoreSuite) SetUpTest(c *check.C) {
	s.Model = &Model{
		Id:   bson.NewObjectId(),
		Age:  22,
		Name: "Mob D.",
	}

	c.Assert(s.Mongo.C("t").Insert(s.Model), check.IsNil)
}

func (s *StoreSuite) TestDialAndClose(c *check.C) {
	store := &Store{}

	c.Check(store.Dial(s.Url), check.IsNil)
	c.Check(store.Dial(s.Url), check.Equals, ErrStillConnected)
	c.Check(store.Close(), check.IsNil)
	c.Check(store.Close(), check.Equals, ErrNotConnected)
}

func (s *StoreSuite) TestFind(c *check.C) {
	m := []Model{}
	q := Query{}

	c.Check(s.Store.Find("t", q, &m), check.IsNil)
	c.Check(m, check.DeepEquals, []Model{*s.Model})
}

func (s *StoreSuite) TestFindOne(c *check.C) {
	m := &Model{}
	q := Query{"_id": s.Model.Id}

	c.Check(s.Store.FindOne("t", q, m), check.IsNil)
	c.Check(m, check.DeepEquals, s.Model)
}

func (s *StoreSuite) TestInsert(c *check.C) {
	m1 := &Model{
		Id:   bson.NewObjectId(),
		Age:  32,
		Name: "Big L.",
	}
	m2 := &Model{}

	c.Check(s.Store.Insert("t", m1), check.IsNil)
	c.Check(s.Mongo.C("t").FindId(m1.Id).One(m2), check.IsNil)
	c.Check(m2, check.DeepEquals, m1)
}

func (s *StoreSuite) TestUpdate(c *check.C) {
	s.Model.Age += 10
	q := Query{"_id": s.Model.Id}
	m := &Model{}

	c.Check(s.Store.Update("t", q, s.Model), check.IsNil)
	c.Check(s.Mongo.C("t").FindId(s.Model.Id).One(m), check.IsNil)
	c.Check(m, check.DeepEquals, s.Model)
}

func (s *StoreSuite) TestRemove(c *check.C) {
	q := Query{"_id": s.Model.Id}

	c.Check(s.Store.Remove("t", q), check.IsNil)
	c.Check(s.Mongo.C("t").FindId(s.Model.Id).One(nil), check.Equals, mgo.ErrNotFound)
}

func (s *StoreSuite) TestRemoveAll(c *check.C) {
	q := Query{}

	c.Check(s.Store.RemoveAll("t", q), check.IsNil)
	c.Check(s.Mongo.C("t").FindId(s.Model.Id).One(nil), check.Equals, mgo.ErrNotFound)
}

func (s *StoreSuite) TearDownTest(c *check.C) {
	c.Assert(s.Mongo.C("t").DropCollection(), check.IsNil)
}

func (s *StoreSuite) TearDownSuite(c *check.C) {
	s.Mongo.Session.Close()
}
