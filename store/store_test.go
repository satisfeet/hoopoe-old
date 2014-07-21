package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestStore(t *testing.T) {
	m := StoreModel{
		Id: bson.NewObjectId(),
	}

	check.Suite(&StoreSuite{
		url:   "localhost/test",
		name:  "testers",
		model: m,
	})
	check.TestingT(t)
}

type StoreModel struct {
	Id   bson.ObjectId `bson:"_id"`
	Text string
}

type StoreSuite struct {
	url   string
	name  string
	mongo *mgo.Session
	model StoreModel
}

func (s *StoreSuite) TestInsert(c *check.C) {
	m := StoreModel{
		Text: "I am getting inserted!!",
	}

	c.Check(NewStore(s.name).Insert(&m), check.IsNil)

	i, err := s.mongo.DB("").C(s.name).Find(nil).Count()

	c.Check(err, check.IsNil)
	c.Check(i, check.Equals, 2)
}

func (s *StoreSuite) TestUpdate(c *check.C) {
	q := Query{"_id": s.model.Id}
	m := StoreModel{}

	s.model.Text += "1234?"

	c.Check(NewStore(s.name).Update(q, &s.model), check.IsNil)

	err := mongo.DB(Database).C(s.name).Find(q).One(&m)

	c.Check(err, check.IsNil)
	c.Check(m, check.DeepEquals, s.model)
}

func (s *StoreSuite) TestRemove(c *check.C) {
	q := Query{"_id": s.model.Id}

	c.Check(NewStore(s.name).Remove(q), check.IsNil)

	i, err := s.mongo.DB("").C(s.name).Find(q).Count()

	c.Check(err, check.IsNil)
	c.Check(i, check.Equals, 0)
}

func (s *StoreSuite) SetUpSuite(c *check.C) {
	var err error
	s.mongo, err = mgo.Dial(s.url)

	c.Assert(err, check.IsNil)
	c.Assert(s.mongo, check.NotNil)
	c.Assert(Open(s.url), check.IsNil)
}

func (s *StoreSuite) TearDownSuite(c *check.C) {
	s.mongo.Close()

	Close()
}

func (s *StoreSuite) SetUpTest(c *check.C) {
	c.Assert(s.mongo.DB("").C(s.name).Insert(s.model), check.IsNil)
}

func (s *StoreSuite) TearDownTest(c *check.C) {
	c.Assert(s.mongo.DB("").C(s.name).DropCollection(), check.IsNil)
}
