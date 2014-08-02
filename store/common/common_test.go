package common

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestCommon(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct{}

type person struct {
	Name    string
	address address
}

type address struct {
	City string
}

func (s *Suite) TestGetFieldValue(c *check.C) {
	p := person{"Bodo", address{"Berlin"}}

	c.Check(GetFieldValue(p, "Name"), check.Equals, "Bodo")
	c.Check(GetFieldValue(&p, "Name"), check.Equals, "Bodo")
}

func (s *Suite) TestSetFieldValue(c *check.C) {
	p := person{}

	SetFieldValue(&p, "Name", "Joe")
	c.Check(p.Name, check.Equals, "Joe")
}

func (s *Suite) TestGetStructInfo(c *check.C) {

}
