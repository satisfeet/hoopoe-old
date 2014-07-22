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
	conf := &Conf{
		Auth:  "foo:bar",
		Host:  "localhost:3000",
		Mongo: "mongodb://localhost/test",
	}
	c.Check(conf.Check(), check.IsNil)
}

func (s *ConfSuite) TestFlags(c *check.C) {
	conf := &Conf{}

	c.Check(conf.Flags([]string{
		"--auth",
		"foo:bar",
		"--host",
		"localhost:3000",
		"--mongo",
		"mongodb://localhost/test",
	}), check.IsNil)

	c.Check(conf.Auth, check.Equals, "foo:bar")
	c.Check(conf.Host, check.Equals, "localhost:3000")
	c.Check(conf.Mongo, check.Equals, "mongodb://localhost/test")
}
