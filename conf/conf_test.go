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
		Username: "foo",
		Password: "bar",
		Host:     "localhost:3000",
		Mongo:    "mongodb://localhost/test",
	}
	c.Check(conf.Check(), check.IsNil)
}

func (s *ConfSuite) TestFlags(c *check.C) {
	conf := &Conf{}

	c.Check(conf.Flags([]string{
		"--username",
		"foo",
		"--password",
		"bar",
		"--host",
		"localhost:3000",
		"--mongo",
		"mongodb://localhost/test",
	}), check.IsNil)

	c.Check(conf.Username, check.Equals, "foo")
	c.Check(conf.Password, check.Equals, "bar")
	c.Check(conf.Host, check.Equals, "localhost:3000")
	c.Check(conf.Mongo, check.Equals, "mongodb://localhost/test")
}
