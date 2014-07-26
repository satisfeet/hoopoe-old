package store

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestStore(t *testing.T) {
	check.Suite(&StoreSuite{
		Url: "localhost/test",
	})
	check.TestingT(t)
}

type StoreSuite struct {
	Url string
}

func (s *StoreSuite) TestOpenAndClose(c *check.C) {
	c.Assert(Open(s.Url), check.IsNil)
	c.Assert(Close(), check.IsNil)
}
