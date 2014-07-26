package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func TestQuery(t *testing.T) {
	check.Suite(&QuerySuite{
		id: bson.NewObjectId(),
	})
	check.TestingT(t)
}

type QuerySuite struct {
	id bson.ObjectId
}

func (s *QuerySuite) TestId(c *check.C) {
	q := Query{}
	c.Check(q.Id(s.id.Hex()), check.IsNil)
	c.Check(q, check.DeepEquals, Query{"_id": s.id})

	q = Query{}
	c.Check(q.Id("1234"), check.Equals, ErrBadIdQuery)
	c.Check(q, check.DeepEquals, Query{})
}

func (s *QuerySuite) TestSearch(c *check.C) {
	f := []string{"foo", "baz"}

	q := Query{}
	c.Check(q.Search("", f), check.Equals, ErrBadSearchQuery)
	c.Check(q, check.DeepEquals, Query{})

	q = Query{}
	c.Check(q.Search("bar", f), check.IsNil)
	c.Check(q, check.DeepEquals, Query{
		"$or": []bson.M{
			bson.M{"foo": bson.RegEx{"bar", "i"}},
			bson.M{"baz": bson.RegEx{"bar", "i"}},
		},
	})
}
