package common

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestStore(t *testing.T) {
	check.Suite(&StoreSuite{})
	check.TestingT(t)
}

type Model struct {
	Bar string
}

type StoreSuite struct{}

func (s *StoreSuite) TestName(c *check.C) {
	f := Model{}
	fs := []Model{}

	c.Check(Name(f), check.Equals, "models")
	c.Check(Name(&f), check.Equals, "models")
	c.Check(Name(fs), check.Equals, "models")
	c.Check(Name(&fs), check.Equals, "models")
}
