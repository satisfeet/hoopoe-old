package common

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestStore(t *testing.T) {
	check.Suite(&StoreSuite{})
	check.TestingT(t)
}

type Foo struct {
	Id  FooId
	Bar string
}

func (f Foo) Validate() error {
	return nil
}

type FooId string

func (f FooId) String() string {
	return string(f)
}

type StoreSuite struct{}

func (s *StoreSuite) TestGetId(c *check.C) {
	f := Foo{"1", "hi"}

	c.Check(GetId(f), check.Equals, FooId("1"))
	c.Check(GetId(&f), check.Equals, FooId("1"))
	c.Check(GetId(Foo{}), check.Equals, FooId(""))
}

func (s *StoreSuite) TestGetName(c *check.C) {
	f := Foo{}
	fs := []Foo{}

	c.Check(GetName(f), check.Equals, "foo")
	c.Check(GetName(&f), check.Equals, "foo")
	c.Check(GetName(fs), check.Equals, "foo")
	c.Check(GetName(&fs), check.Equals, "foo")
}
