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

type QuerySuite struct{}

func (s *QuerySuite) TestId(c *check.C) {
	q := Query{}
	c.Check(q.Id(0), check.Equals, ErrBadQueryParam)
	c.Check(q.Id(nil), check.Equals, ErrBadQueryParam)
	c.Check(q.Id("abcd"), check.Equals, ErrBadQueryParam)
	c.Check(q.Id(bson.ObjectId("abcd")), check.Equals, ErrBadQueryParam)

	id := bson.NewObjectId()
	q = Query{}
	c.Check(q.Id(id.Hex()), check.IsNil)
	c.Check(q["_id"], check.Equals, id)
	q = Query{}
	c.Check(q.Id(id), check.IsNil)
	c.Check(q["_id"], check.Equals, id)
}

func (s *QuerySuite) TestOr(c *check.C) {
	q := Query{}
	c.Check(q.Or(Query{"child": 1}), check.IsNil)
	c.Check(q.Or(Query{"child": 2}), check.IsNil)
	c.Check(q["$or"], check.DeepEquals, []Query{
		Query{"child": 1},
		Query{"child": 2},
	})
}

func (s *QuerySuite) TestRegex(c *check.C) {
	q := Query{}
	c.Check(q.Regex("foo", "bar"), check.IsNil)
	c.Check(q["foo"], check.Equals, bson.RegEx{"bar", "i"})
}
