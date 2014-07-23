package validation

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestError(t *testing.T) {
	check.Suite(&ErrorSuite{})
	check.TestingT(t)
}

type ErrorSuite struct{}

func (s *ErrorSuite) TestHas(c *check.C) {
	c.Check(Error{}.Has(), check.Equals, false)

	c.Check(Error{
		"foo": ErrRequired,
	}.Has(), check.Equals, true)
	c.Check(Error{
		"foo": ErrRequired,
		"bar": ErrRange,
	}.Has(), check.Equals, true)

	err := Error{}
	err.Set("baz", ErrEmail)
	c.Check(err.Has(), check.Equals, true)
}

func (s *ErrorSuite) TestSet(c *check.C) {
	err := Error{}

	err.Set("foo", nil)
	err.Set("bar", ErrLength)

	c.Check(err["foo"], check.IsNil)
	c.Check(err["bar"], check.Equals, ErrLength)
}

func (s *ErrorSuite) TestError(c *check.C) {
	c.Check(Error{}.Error(), check.Equals, "")
	c.Check(Error{
		"foo": ErrEmail,
	}.Error(), check.Matches, "*invalid email*")
}
