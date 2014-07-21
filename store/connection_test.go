package store

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestConnection(t *testing.T) {
	check.Suite(&ConnectionSuite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type ConnectionSuite struct {
	url string
}

func (s *ConnectionSuite) TestOpen(c *check.C) {
	c.Check(Open(s.url), check.IsNil)
	c.Check(mongo, check.NotNil)
}

func (s *ConnectionSuite) TestClose(c *check.C) {
	Open(s.url)
	Close()
}
