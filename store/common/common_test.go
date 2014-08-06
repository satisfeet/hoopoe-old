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
	Name    string `store:"unique"`
	addresS addresS
}

// We need an upper case letter in name to check lower casing.
type addresS struct {
	City string `store:"index"`
}

func (s *Suite) TestGetTypeName(c *check.C) {
	p := person{}
	ps := []person{p}

	c.Check(GetTypeName(p), check.Equals, "person")
	c.Check(GetTypeName(&p), check.Equals, "person")
	c.Check(GetTypeName(ps), check.Equals, "person")
	c.Check(GetTypeName(&ps), check.Equals, "person")
}

func (s *Suite) TestGetFieldValue(c *check.C) {
	p := person{"Bodo", addresS{"Berlin"}}

	c.Check(GetFieldValue(p, "Name"), check.Equals, "Bodo")
	c.Check(GetFieldValue(&p, "Name"), check.Equals, "Bodo")
}

func (s *Suite) TestSetFieldValue(c *check.C) {
	p := person{}

	SetFieldValue(&p, "Name", "Joe")
	c.Check(p.Name, check.Equals, "Joe")
}

func (s *Suite) TestGetStructInfo(c *check.C) {
	result := GetStructInfo(person{})

	c.Check(result, check.DeepEquals, map[string]FieldInfo{
		"name":         FieldInfo{Name: "Name", Unique: true},
		"address.city": FieldInfo{Name: "City", Index: true},
	})
}
