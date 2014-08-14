package mongo

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func TestTypes(t *testing.T) {
	check.Suite(&TypesSuite{})
	check.TestingT(t)
}

type TypesSuite struct{}

func (s *TypesSuite) TestIdFromString(c *check.C) {
	id := bson.NewObjectId()

	c.Check(IdFromString(id.Hex()).Valid(), check.Equals, true)
	c.Check(IdFromString("123567").Valid(), check.Equals, false)
}
