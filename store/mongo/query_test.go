package mongo

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func TestQuery(t *testing.T) {
	check.Suite(&QuerySuite{})
	check.TestingT(t)
}

type QuerySuite struct {
	id bson.ObjectId
}

func (s *QuerySuite) SetUpTest(c *check.C) {
	s.id = bson.NewObjectId()
}

func (s *QuerySuite) TestId(c *check.C) {
	q1 := Query{}
	q2 := Query{}
	q3 := Query{}

	q1.Id(s.id)
	q2.Id(s.id.Hex())
	q3.Id("abcd")

	c.Check(q1["_id"], check.Equals, s.id)
	c.Check(q2["_id"], check.Equals, s.id)
	c.Check(q3["_id"], check.IsNil)
}

func (s *QuerySuite) TestIn(c *check.C) {
	q := Query{}
	q.In("foo", 123)

	c.Check(q["foo"], check.DeepEquals, bson.M{
		"$in": []interface{}{123},
	})
}

func (s *QuerySuite) TestPush(c *check.C) {
	u := Query{}
	u.Push("foo", "bar")

	c.Check(u["$push"], check.DeepEquals, bson.M{"foo": "bar"})
}

func (s *QuerySuite) TestPull(c *check.C) {
	u := Query{}
	u.Pull("foo", "bar")

	c.Check(u["$pull"], check.DeepEquals, bson.M{"foo": "bar"})
}
