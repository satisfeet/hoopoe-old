package store

import (
	"testing"

	"gopkg.in/check.v1"

	"github.com/satisfeet/hoopoe/store/mongo"
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

func (s *StoreSuite) TestDialAndClose(c *check.C) {
	c.Check(Dial(s.Url), check.IsNil)
	c.Check(Dial(s.Url), check.Equals, mongo.ErrStillConnected)
	c.Check(Close(), check.IsNil)
	c.Check(Close(), check.Equals, mongo.ErrNotConnected)
}
