package validation

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/validator.v1"
)

func TestRules(t *testing.T) {
	check.Suite(&RuleSuite{})
	check.TestingT(t)
}

type RuleSuite struct{}

func (s *RuleSuite) TestMin(c *check.C) {
	c.Check(Min(5, "4"), check.IsNil)
	c.Check(Min("abc", "2"), check.IsNil)

	c.Check(Min(10, "11"), check.Equals, validator.ErrMin)
	c.Check(Min("abc", "4"), check.Equals, validator.ErrMin)

	c.Check(Min(1, "ab"), check.Equals, validator.ErrBadParameter)
	c.Check(Min(true, "12"), check.Equals, validator.ErrUnsupported)
}

func (s *RuleSuite) TestEmail(c *check.C) {
	c.Check(Email("i@foo.me", ""), check.IsNil)

	c.Check(Email(1234, ""), check.Equals, validator.ErrUnsupported)
	c.Check(Email("@foobar.me", ""), check.Equals, validator.ErrInvalid)
}
