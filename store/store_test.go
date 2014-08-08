package store

import (
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *Suite) TestParseId(c *check.C) {
	id := bson.NewObjectId()

	c.Check(ParseId(id).Valid(), check.Equals, true)
	c.Check(ParseId(id.Hex()).Valid(), check.Equals, true)

	c.Check(ParseId(1).Valid(), check.Equals, false)
	c.Check(ParseId(nil).Valid(), check.Equals, false)
	c.Check(ParseId("123").Valid(), check.Equals, false)
}
