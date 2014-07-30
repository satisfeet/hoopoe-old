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
	id1 := bson.NewObjectId()
	id2 := bson.ObjectId("abcd")

	c.Check(q.Id(0), check.Equals, ErrBadQueryParam)
	c.Check(q.Id(nil), check.Equals, ErrBadQueryParam)
	c.Check(q.Id(id2), check.Equals, ErrBadQueryParam)
	c.Check(q.Id("abcd"), check.Equals, ErrBadQueryParam)

	q = Query{}
	c.Check(q.Id(id1.Hex()), check.IsNil)
	q = Query{}
	c.Check(q.Id(id1), check.IsNil)
	c.Check(q["_id"], check.Equals, id1)
}

func (s *QuerySuite) TestOr(c *check.C) {
	q := Query{}
	q1 := Query{"child": 1}
	q2 := Query{"child": 2}

	c.Check(q.Or(q1), check.IsNil)
	c.Check(q.Or(q2), check.IsNil)
	c.Check(q["$or"], check.DeepEquals, []Query{q1, q2})
}

func (s *QuerySuite) TestRegex(c *check.C) {
	q := Query{}

	c.Check(q.Regex("foo", "bar"), check.IsNil)
	c.Check(q["foo"], check.Equals, bson.RegEx{"bar", "i"})
}
