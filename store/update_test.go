package store

import "gopkg.in/check.v1"

func (s *Suite) TestUpdatePushId(c *check.C) {
	u := Update{}

	c.Check(u.PushId("foo", nil), check.Equals, ErrBadParam)
	c.Check(u["$push"], check.IsNil)
	c.Check(u.PushId("foo", s.id), check.IsNil)
	c.Check(u["$push"], check.NotNil)
}

func (s *Suite) TestUpdatePullId(c *check.C) {
	u := Update{}

	c.Check(u.PullId("foo", nil), check.Equals, ErrBadParam)
	c.Check(u["$pull"], check.IsNil)
	c.Check(u.PullId("foo", s.id), check.IsNil)
	c.Check(u["$pull"], check.NotNil)
}
