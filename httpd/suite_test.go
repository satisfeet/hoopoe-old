package httpd

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
)

func TestHttpd(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	url string
	db  *mgo.Database
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.db = sess.DB("")
}

func (s *Suite) TearDownSuite(c *check.C) {
	s.db.Session.Close()
}
