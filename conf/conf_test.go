package conf

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestConf(t *testing.T) {
	check.Suite(&ConfSuite{
		Conf: &Conf{},
	})
	check.TestingT(t)
}

type ConfSuite struct {
	Conf *Conf
}

func (s *ConfSuite) TestCheck(c *check.C) {
	s.Conf.Username = "foo"
	s.Conf.Password = "bar"
	s.Conf.Host = "localhost:3000"
	s.Conf.Mongo = "mongodb://localhost/test"

	c.Check(s.Conf.Check(), check.IsNil)
}

func (s *ConfSuite) TestFlags(c *check.C) {
	c.Check(s.Conf.Flags([]string{
		"--username",
		"foo",
		"--password",
		"bar",
		"--host",
		"localhost:3000",
		"--mongo",
		"mongodb://localhost/test",
	}), check.IsNil)

	c.Check(s.Conf.Username, check.Equals, "foo")
	c.Check(s.Conf.Password, check.Equals, "bar")
	c.Check(s.Conf.Host, check.Equals, "localhost:3000")
	c.Check(s.Conf.Mongo, check.Equals, "mongodb://localhost/test")
}
