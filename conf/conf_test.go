package conf

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestConf(t *testing.T) {
	check.Suite(&ConfSuite{})
	check.TestingT(t)
}

type ConfSuite struct{}

func (s *ConfSuite) TestCheck(c *check.C) {
	Username = "foo"
	Password = "bar"
	Host = "localhost:3000"
	Mongo = "mongodb://localhost/test"

	c.Check(Check(), check.IsNil)
}

func (s *ConfSuite) TestFlags(c *check.C) {
	c.Check(Flags([]string{
		"--username",
		"foo",
		"--password",
		"bar",
		"--host",
		"localhost:3000",
		"--mongo",
		"mongodb://localhost/test",
	}), check.IsNil)

	c.Check(Username, check.Equals, "foo")
	c.Check(Password, check.Equals, "bar")
	c.Check(Host, check.Equals, "localhost:3000")
	c.Check(Mongo, check.Equals, "mongodb://localhost/test")
}
