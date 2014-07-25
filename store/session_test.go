package store

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestSession(t *testing.T) {
	check.Suite(&SessionSuite{
		url:     "localhost/test",
		session: &Session{},
	})
	check.TestingT(t)
}

type SessionSuite struct {
	url     string
	session *Session
}

func (s *SessionSuite) TestOpen(c *check.C) {
	c.Check(s.session.Open(s.url), check.IsNil)

	m, err := s.session.Mongo()

	c.Assert(err, check.IsNil)

	m.Close()
}

func (s *SessionSuite) TestClose(c *check.C) {
	s.session.Open(s.url)

	c.Assert(s.session.Close(), check.IsNil)
}
