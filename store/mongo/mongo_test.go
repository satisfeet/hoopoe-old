package mongo

import (
	"testing"

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

type model struct {
	Id     bson.ObjectId `bson:"_id"`
	Age    int           `store:"index"`
	Name   string        `store:"unique"`
	Nested nested
}

type nested struct {
	Foo string `store:"index"`
}

type Suite struct {
	url        string
	model      model
	store      *Store
	database   *mgo.Database
	collection *mgo.Collection
}

func (s *Suite) TestQueryId(c *check.C) {
	q := Query{}
	c.Check(q.Id(0), check.Equals, ErrBadQueryParam)
	c.Check(q.Id(nil), check.Equals, ErrBadQueryParam)
	c.Check(q.Id("abcd"), check.Equals, ErrBadQueryParam)
	c.Check(q.Id(bson.ObjectId("abcd")), check.Equals, ErrBadQueryParam)

	id := bson.NewObjectId()
	q = Query{}
	c.Check(q.Id(id.Hex()), check.IsNil)
	c.Check(q["_id"], check.Equals, id)
	q = Query{}
	c.Check(q.Id(id), check.IsNil)
	c.Check(q["_id"], check.Equals, id)
}

func (s *Suite) TestQueryOr(c *check.C) {
	q := Query{}
	c.Check(q.Or(Query{"child": 1}), check.IsNil)
	c.Check(q.Or(Query{"child": 2}), check.IsNil)
	c.Check(q["$or"], check.DeepEquals, []Query{
		Query{"child": 1},
		Query{"child": 2},
	})
}

func (s *Suite) TestQueryRegex(c *check.C) {
	q := Query{}
	c.Check(q.Regex("foo", "bar"), check.IsNil)
	c.Check(q["foo"], check.Equals, bson.RegEx{"bar", "i"})
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.store = &Store{
		session:  sess,
		database: sess.DB(""),
	}
	s.database = sess.DB("")
	s.collection = sess.DB("").C("models")
}

func (s *Suite) SetUpTest(c *check.C) {
	s.model = model{
		Id:   bson.NewObjectId(),
		Age:  22,
		Name: "Mob D.",
	}

	err := s.collection.Insert(s.model)
	c.Assert(err, check.IsNil)
}

func (s *Suite) TestStoreDialAndClose(c *check.C) {
	store := Store{}

	c.Assert(store.Dial(s.url), check.IsNil)
	c.Assert(store.Dial(s.url), check.Equals, ErrStillConnected)
	c.Assert(store.Close(), check.IsNil)
	c.Assert(store.Close(), check.Equals, ErrNotConnected)
}

func (s *Suite) TestStoreFind(c *check.C) {
	m := []model{}

	err := s.store.Find(Query{}, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, []model{s.model})
}

func (s *Suite) TestStoreFindOne(c *check.C) {
	m := model{}

	err := s.store.FindOne(Query{"_id": s.model.Id}, &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, s.model)
}

func (s *Suite) TestStoreIndex(c *check.C) {
	err := s.store.Index(s.model)
	c.Assert(err, check.IsNil)

	indexes, err := s.collection.Indexes()
	c.Assert(err, check.IsNil)

	hasIndex := false
	hasUnique := false

	for _, i := range indexes {
		switch i.Name {
		case "name_1":
			hasUnique = true

			c.Check(i.Key, check.DeepEquals, []string{"name"})
			c.Check(i.Unique, check.Equals, true)
		case "nested.foo_1_age_1":
			hasIndex = true

			c.Check(i.Key, check.DeepEquals, []string{"nested.foo", "age"})
			c.Check(i.Unique, check.Equals, false)
		case "age_1_nested.foo_1":
			hasIndex = true

			c.Check(i.Key, check.DeepEquals, []string{"age", "nested.foo"})
			c.Check(i.Unique, check.Equals, false)
		}
	}

	c.Check(hasIndex, check.Equals, true)
	c.Check(hasUnique, check.Equals, true)
}

func (s *Suite) TestStoreInsert(c *check.C) {
	m1 := model{
		Age:  32,
		Name: "Big L.",
	}
	m2 := model{}

	err := s.store.Insert(&m1)
	c.Assert(err, check.IsNil)

	err = s.collection.FindId(m1.Id).One(&m2)
	c.Assert(err, check.IsNil)

	c.Check(m2, check.DeepEquals, m1)
}

func (s *Suite) TestStoreUpdate(c *check.C) {
	s.model.Age += 10
	m := model{}

	err := s.store.Update(s.model)
	c.Assert(err, check.IsNil)

	err = s.collection.FindId(s.model.Id).One(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, s.model)
}

func (s *Suite) TestStoreRemove(c *check.C) {
	err := s.store.Remove(s.model)
	c.Assert(err, check.IsNil)

	err = s.collection.FindId(s.model.Id).One(nil)
	c.Check(err, check.Equals, mgo.ErrNotFound)
}

func (s *Suite) TearDownTest(c *check.C) {
	_, err := s.collection.RemoveAll(nil)
	c.Assert(err, check.IsNil)

	indexes, err := s.collection.Indexes()
	c.Assert(err, check.IsNil)

	for _, index := range indexes {
		if index.Key[0] != "_id" {
			err = s.collection.DropIndex(index.Key...)
			c.Assert(err, check.IsNil)
		}
	}
}

func (s *Suite) TearDownSuite(c *check.C) {
	s.database.Session.Close()
}
