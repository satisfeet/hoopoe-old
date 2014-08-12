package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

type model struct {
	Name string `store:"index"`
}

func TestSuite(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct {
	id bson.ObjectId
}

func (s *Suite) SetUpTest(c *check.C) {
	s.id = bson.NewObjectId()
}

func (s *Suite) TestQueryId(c *check.C) {
	q := query{}
	q.Id(s.id)

	c.Check(q["_id"], check.Equals, s.id)
}

func (s *Suite) TestQueryIn(c *check.C) {
	q := query{}
	q.In("foo", 123)

	c.Check(q["foo"], check.DeepEquals, bson.M{
		"$in": []interface{}{123},
	})
}

func (s *Suite) TestQueryPush(c *check.C) {
	u := query{}
	u.Push("foo", "bar")

	c.Check(u["$push"], check.DeepEquals, bson.M{"foo": "bar"})
}

func (s *Suite) TestQueryPull(c *check.C) {
	u := query{}
	u.Pull("foo", "bar")

	c.Check(u["$pull"], check.DeepEquals, bson.M{"foo": "bar"})
}

func (s *Suite) TestParseId(c *check.C) {
	id := bson.NewObjectId()

	c.Check(ParseId(id).Valid(), check.Equals, true)
	c.Check(ParseId(id.Hex()).Valid(), check.Equals, true)

	c.Check(ParseId(1).Valid(), check.Equals, false)
	c.Check(ParseId(nil).Valid(), check.Equals, false)
	c.Check(ParseId("123").Valid(), check.Equals, false)
}
