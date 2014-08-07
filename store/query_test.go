package store

import (
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func (s *Suite) TestQueryId(c *check.C) {
	q := Query{}

	c.Check(q.Id(1), check.Equals, ErrBadParam)
	c.Check(q.Id(nil), check.Equals, ErrBadParam)
	c.Check(q.Id("123"), check.Equals, ErrBadParam)

	id := bson.NewObjectId()

	q = Query{}
	c.Check(q.Id(id), check.IsNil)
	c.Check(q["_id"], check.Equals, id)
	q = Query{}
	c.Check(q.Id(id.Hex()), check.IsNil)
	c.Check(q["_id"], check.Equals, id)
}

func (s *Suite) TestQuerySearch(c *check.C) {
	q := Query{}

	c.Check(q.Search("", model{}), check.Equals, ErrBadParam)

	c.Check(q.Search("foo", model{}), check.IsNil)
	c.Check(q["$or"], check.DeepEquals, []bson.M{
		bson.M{"name": bson.RegEx{"foo", "i"}},
	})
}
