package validation

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestValidation(t *testing.T) {
	check.Suite(&ValidationSuite{})
	check.TestingT(t)
}

type ValidationSuite struct{}

func (s *ValidationSuite) TestEmail(c *check.C) {
	c.Check(Email("foo@bar.com"), check.IsNil)
	c.Check(Email("foo"), check.Equals, ErrEmail)
	c.Check(Email("foo@"), check.Equals, ErrEmail)
	c.Check(Email("foo@bar."), check.Equals, ErrEmail)
}

func (s *ValidationSuite) TestRange(c *check.C) {
	c.Check(Range(2, 1, 3), check.IsNil)
	c.Check(Range(0, 1, 0), check.Equals, ErrRange)
	c.Check(Range(11, 0, 10), check.Equals, ErrRange)
}

func (s *ValidationSuite) TestLength(c *check.C) {
	c.Check(Length("ab", 1, 3), check.IsNil)
	c.Check(Length("", 1, 0), check.Equals, ErrLength)
	c.Check(Length("abc", 0, 2), check.Equals, ErrLength)
}

func (s *ValidationSuite) TestRequired(c *check.C) {
	var a int
	var b string

	c.Check(Required(nil), check.Equals, ErrRequired)
	c.Check(Required(""), check.Equals, ErrRequired)
	c.Check(Required(0), check.Equals, ErrRequired)
	c.Check(Required(a), check.Equals, ErrRequired)
	c.Check(Required(b), check.Equals, ErrRequired)
	c.Check(Required("ab"), check.IsNil)
}
