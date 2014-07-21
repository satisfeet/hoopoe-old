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
	q.Id(s.id.Hex())
	c.Check(q, check.DeepEquals, Query{"_id": s.id})

	q = Query{}
	q.Id("1234")
	c.Check(q, check.DeepEquals, Query{})
}

func (s *QuerySuite) TestValid(c *check.C) {
	c.Check(Query{"_id": s.id}.Valid(), check.Equals, true)
	c.Check(Query{}.Valid(), check.Equals, false)
}

func (s *QuerySuite) TestSearch(c *check.C) {
	f := []string{"foo", "baz"}

	q := Query{}
	q.Search("", f)
	c.Check(q, check.DeepEquals, Query{})

	q = Query{}
	q.Search("bar", f)
	c.Check(q, check.DeepEquals, Query{
		"$or": []bson.M{
			bson.M{"foo": bson.RegEx{"bar", "i"}},
			bson.M{"baz": bson.RegEx{"bar", "i"}},
		},
	})
}
