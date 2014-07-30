package store

import (
	"testing"

	"github.com/satisfeet/hoopoe/store/mongo"
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
	Id    bson.ObjectId `bson:"_id"`
	Name  string        `store:"index"`
	Email string        `store:"unique"`
}

type StoreSuite struct {
	Url        string
	Model      *StoreModel
	Store      *Store
	Database   *mgo.Database
	Collection *mgo.Collection
}

func (suite *StoreSuite) SetUpSuite(c *check.C) {
	s, err := mgo.Dial(suite.Url)
	c.Assert(err, check.IsNil)

	m := new(mongo.Store)
	c.Assert(m.Dial(suite.Url), check.IsNil)

	suite.Store = &Store{m}
	suite.Database = s.DB("")
	suite.Collection = s.DB("").C("storemodels")
}

func (suite *StoreSuite) SetUpTest(c *check.C) {
	suite.Model = &StoreModel{
		Id:    bson.NewObjectId(),
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
	}
	c.Assert(suite.Collection.Insert(suite.Model), check.IsNil)
}

func (suite *StoreSuite) TestDialAndClose(c *check.C) {
	s := NewStore()

	c.Assert(s.Dial(suite.Url), check.IsNil)
	c.Assert(s.Dial(suite.Url), check.NotNil)
	c.Assert(s.Close(), check.IsNil)
	c.Assert(s.Close(), check.NotNil)
}

func (suite *StoreSuite) TestInsert(c *check.C) {
	m1 := &StoreModel{
		Id:    bson.NewObjectId(),
		Name:  "Jason Statham",
		Email: "j@stats.uk",
	}
	m2 := new(StoreModel)

	c.Assert(suite.Store.Insert(m1), check.IsNil)
	c.Check(suite.Collection.FindId(m1.Id).One(m2), check.IsNil)
	c.Check(m1, check.DeepEquals, m2)
}

func (suite *StoreSuite) TestUpdate(c *check.C) {
	m := new(StoreModel)
	suite.Model.Name += " I"

	c.Assert(suite.Store.Update(suite.Model), check.IsNil)
	c.Check(suite.Collection.FindId(suite.Model.Id).One(m), check.IsNil)
	c.Check(m, check.DeepEquals, suite.Model)
}

func (suite *StoreSuite) TestRemove(c *check.C) {
	c.Assert(suite.Store.Remove(suite.Model), check.IsNil)
	c.Check(suite.Collection.FindId(suite.Model.Id).One(nil), check.Equals, mgo.ErrNotFound)
}

func (suite *StoreSuite) TearDownTest(c *check.C) {
	c.Assert(suite.Collection.DropCollection(), check.IsNil)
}

func (suite *StoreSuite) TearDownSuite(c *check.C) {
	suite.Database.Session.Close()
}
