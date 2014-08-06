package utils

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestUtils(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct{}

type Person struct {
	Name    string `store:"unique"`
	Address Address
}

type Address struct {
	City string `store:"index"`
}

func (s *Suite) TestGetTypeName(c *check.C) {
	p := Person{}
	ps := []Person{p}

	c.Check(GetTypeName(p), check.Equals, "Person")
	c.Check(GetTypeName(&p), check.Equals, "Person")
	c.Check(GetTypeName(ps), check.Equals, "Person")
	c.Check(GetTypeName(&ps), check.Equals, "Person")
}

func (s *Suite) TestGetFieldValue(c *check.C) {
	p := Person{"Bodo", Address{"Berlin"}}

	c.Check(GetFieldValue(p, "Name"), check.Equals, "Bodo")
	c.Check(GetFieldValue(&p, "Name"), check.Equals, "Bodo")
}

func (s *Suite) TestSetFieldValue(c *check.C) {
	p := Person{}

	SetFieldValue(&p, "Name", "Joe")
	c.Check(p.Name, check.Equals, "Joe")
}

func (s *Suite) TestGetFieldValues(c *check.C) {
	p1 := Person{"Bodo", Address{"Berlin"}}
	p2 := Person{}

	c.Check(GetFieldValues(p1), check.DeepEquals, map[string]interface{}{
		"name": "Bodo",
		"address": map[string]interface{}{
			"city": "Berlin",
		},
	})
	c.Check(GetFieldValues(p2), check.DeepEquals, map[string]interface{}{})
}

func (s *Suite) TestGetStructInfo(c *check.C) {
	result := GetStructInfo(Person{})

	c.Check(result, check.DeepEquals, map[string]FieldInfo{
		"name":         FieldInfo{Name: "Name", Unique: true},
		"address.city": FieldInfo{Name: "City", Index: true},
	})
}
