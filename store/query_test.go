package store

import (
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *Suite) TestQueryId(c *check.C) {
	q := Query{}

	c.Check(q.Id(nil), check.Equals, ErrBadParam)
	c.Check(q["_id"], check.IsNil)
	c.Check(q.Id(s.id), check.IsNil)
	c.Check(q["_id"], check.Equals, s.id)
}

func (s *Suite) TestQueryHasId(c *check.C) {
	q := Query{}

	c.Check(q.HasId("foo", nil), check.Equals, ErrBadParam)
	c.Check(q["foo"], check.IsNil)
	c.Check(q.HasId("foo", s.id), check.IsNil)
	c.Check(q["foo"], check.NotNil)
}

func (s *Suite) TestQuerySearch(c *check.C) {
	q := Query{}

	c.Check(q.Search("", model{}), check.Equals, ErrBadParam)

	c.Check(q.Search("foo", model{}), check.IsNil)
	c.Check(q["$or"], check.DeepEquals, []bson.M{
		bson.M{"name": bson.RegEx{"foo", "i"}},
	})
}
