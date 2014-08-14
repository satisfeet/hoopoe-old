package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func TestStore(t *testing.T) {
	check.Suite(&StoreSuite{})
	check.TestingT(t)
}

type StoreSuite struct {
	id bson.ObjectId
}

func (s *StoreSuite) SetUpTest(c *check.C) {
	s.id = bson.NewObjectId()
}

func (s *StoreSuite) TestIdFromString(c *check.C) {
	id1 := IdFromString(s.id.Hex())
	id2 := IdFromString("")

	c.Check(id1.Valid(), check.Equals, true)
	c.Check(id2.Valid(), check.Equals, false)
}
