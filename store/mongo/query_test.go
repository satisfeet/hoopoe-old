package mongo

import (
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *Suite) TestQueryId(c *check.C) {
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

func (s *Suite) TestQueryIn(c *check.C) {
	q := Query{}
	q.In("foo", 123)

	c.Check(q["foo"], check.DeepEquals, bson.M{
		"$in": []interface{}{123},
	})
}

func (s *Suite) TestQueryPush(c *check.C) {
	u := Query{}
	u.Push("foo", "bar")

	c.Check(u["$push"], check.DeepEquals, bson.M{"foo": "bar"})
}

func (s *Suite) TestQueryPull(c *check.C) {
	u := Query{}
	u.Pull("foo", "bar")

	c.Check(u["$pull"], check.DeepEquals, bson.M{"foo": "bar"})
}
