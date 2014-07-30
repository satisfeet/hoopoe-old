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

type StoreModel struct {
	Id   bson.ObjectId `bson:"_id"`
	Age  int
	Name string
}

type StoreSuite struct {
	Url        string
	Store      *Store
	Model      *StoreModel
	Database   *mgo.Database
	Collection *mgo.Collection
}

func (suite *StoreSuite) SetUpSuite(c *check.C) {
	s, err := mgo.Dial(suite.Url)
	c.Assert(err, check.IsNil)
	suite.Store = &Store{
		session:  s,
		database: s.DB(""),
	}
	suite.Database = s.DB("")
	suite.Collection = s.DB("").C("storemodels")
}

func (suite *StoreSuite) SetUpTest(c *check.C) {
	suite.Model = &StoreModel{
		Id:   bson.NewObjectId(),
		Age:  22,
		Name: "Mob D.",
	}
	c.Assert(suite.Collection.Insert(suite.Model), check.IsNil)
}

func (suite *StoreSuite) TestDialAndClose(c *check.C) {
	s := new(Store)
	c.Check(s.Dial(suite.Url), check.IsNil)
	c.Check(s.Dial(suite.Url), check.Equals, ErrStillConnected)
	c.Check(s.Close(), check.IsNil)
	c.Check(s.Close(), check.Equals, ErrNotConnected)
}

func (suite *StoreSuite) TestFind(c *check.C) {
	m := []StoreModel{}
	c.Check(suite.Store.Find(Query{}, &m), check.IsNil)
	c.Check(m, check.DeepEquals, []StoreModel{*suite.Model})
}

func (suite *StoreSuite) TestFindOne(c *check.C) {
	m := &StoreModel{}
	c.Check(suite.Store.FindOne(Query{"_id": suite.Model.Id}, m), check.IsNil)
	c.Check(m, check.DeepEquals, suite.Model)
}

func (suite *StoreSuite) TestInsert(c *check.C) {
	m1 := &StoreModel{
		Id:   bson.NewObjectId(),
		Age:  32,
		Name: "Big L.",
	}
	m2 := new(StoreModel)

	c.Check(suite.Store.Insert(m1), check.IsNil)
	c.Check(suite.Collection.FindId(m1.Id).One(m2), check.IsNil)
	c.Check(m2, check.DeepEquals, m1)
}

func (suite *StoreSuite) TestUpdate(c *check.C) {
	suite.Model.Age += 10
	m := new(StoreModel)

	c.Check(suite.Store.Update(suite.Model), check.IsNil)
	c.Check(suite.Collection.FindId(suite.Model.Id).One(m), check.IsNil)
	c.Check(m, check.DeepEquals, suite.Model)
}

func (suite *StoreSuite) TestRemove(c *check.C) {
	c.Check(suite.Store.Remove(suite.Model), check.IsNil)
	c.Check(suite.Collection.FindId(suite.Model.Id).One(nil), check.Equals, mgo.ErrNotFound)
}

func (suite *StoreSuite) TearDownTest(c *check.C) {
	c.Assert(suite.Collection.DropCollection(), check.IsNil)
}

func (suite *StoreSuite) TearDownSuite(c *check.C) {
	suite.Database.Session.Close()
}
