package mongo

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/common"
)

func TestMongo(t *testing.T) {
	check.Suite(&QuerySuite{})
	check.Suite(&StoreSuite{
		Url: "localhost/test",
	})
	check.TestingT(t)
}

type QuerySuite struct{}

func (s *QuerySuite) TestId(c *check.C) {
	q := Query{}
	id1 := bson.NewObjectId()
	id2 := bson.ObjectId("abcd")

	c.Check(q.Id(0), check.Equals, common.ErrBadQueryValue)
	c.Check(q.Id(nil), check.Equals, common.ErrBadQueryValue)
	c.Check(q.Id(id2), check.Equals, common.ErrBadQueryId)
	c.Check(q.Id("abcd"), check.Equals, common.ErrBadQueryId)

	q = Query{}
	c.Check(q.Id(id1.Hex()), check.IsNil)
	q = Query{}
	c.Check(q.Id(id1), check.IsNil)
	c.Check(q["_id"], check.Equals, id1)
}

func (s *QuerySuite) TestOr(c *check.C) {
	q := Query{}
	q1 := Query{"child": 1}
	q2 := Query{"child": 2}

	c.Check(q.Or(q1), check.IsNil)
	c.Check(q.Or(q2), check.IsNil)
	c.Check(q["$or"], check.DeepEquals, []Query{q1, q2})
}

func (s *QuerySuite) TestRegex(c *check.C) {
	q := Query{}

	c.Check(q.Regex("foo", "bar"), check.IsNil)
	c.Check(q["foo"], check.Equals, bson.RegEx{"bar", "i"})
}

// Test model struct.
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

// Establishes connection to database. Sets up store instance.
func (s *StoreSuite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.Url)
	c.Assert(err, check.IsNil)

	s.Mongo = sess.DB("")
	s.Store = &Store{sess}
}

// Sets up a new model and inserts is into database.
func (s *StoreSuite) SetUpTest(c *check.C) {
	s.Model = &Model{
		Id:   bson.NewObjectId(),
		Age:  22,
		Name: "Mob D.",
	}

	c.Assert(s.Mongo.C("t").Insert(s.Model), check.IsNil)
}

func (s *StoreSuite) TestOpenAndClose(c *check.C) {
	store := &Store{}

	c.Check(store.Open(s.Url), check.IsNil)
	c.Check(store.Open(s.Url), check.Equals, common.ErrStillConnected)
	c.Check(store.Close(), check.IsNil)
	c.Check(store.Close(), check.Equals, common.ErrNotConnected)
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

func (s *StoreSuite) TestFindAll(c *check.C) {
	m := []Model{}
	q := Query{}

	c.Check(s.Store.FindAll("t", q, &m), check.IsNil)
	c.Check(m, check.DeepEquals, []Model{*s.Model})
}

func (s *StoreSuite) TestFindOne(c *check.C) {
	m := &Model{}
	q := Query{"_id": s.Model.Id}

	c.Check(s.Store.FindOne("t", q, m), check.IsNil)
	c.Check(m, check.DeepEquals, s.Model)
}

// Drops database collection for clean up.
func (s *StoreSuite) TearDownTest(c *check.C) {
	c.Assert(s.Mongo.C("t").DropCollection(), check.IsNil)
}

// Closes connections to database.
func (s *StoreSuite) TearDownSuite(c *check.C) {
	s.Mongo.Session.Close()
}
