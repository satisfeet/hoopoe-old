package utils

import (
	"testing"

	"gopkg.in/check.v1"
)

var people = []Person{
	Person{},
	Person{"Bodo", Address{"Berlin"}},
}

type Person struct {
	Name    string `store:"unique"`
	Address Address
}

type Address struct {
	City string `store:"index"`
}

func TestUtils(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct{}

func (s *Suite) TestGetTypeName(c *check.C) {
	c.Check(GetTypeName(people[0]), check.Equals, "Person")
	c.Check(GetTypeName(&people[0]), check.Equals, "Person")
	c.Check(GetTypeName(people), check.Equals, "Person")
	c.Check(GetTypeName(&people), check.Equals, "Person")
}

func (s *Suite) TestGetFieldValue(c *check.C) {
	c.Check(GetFieldValue(people[1], "Name"), check.Equals, "Bodo")
	c.Check(GetFieldValue(&people[1], "Name"), check.Equals, "Bodo")
}

func (s *Suite) TestSetFieldValue(c *check.C) {
	SetFieldValue(&people[0], "Name", "Joe")

	c.Check(people[0].Name, check.Equals, "Joe")
}

func (s *Suite) TestGetFieldValues(c *check.C) {
	c.Check(GetFieldValues(people[0]), check.DeepEquals, map[string]interface{}{})
	c.Check(GetFieldValues(people[1]), check.DeepEquals, map[string]interface{}{
		"name": "Bodo",
		"address": map[string]interface{}{
			"city": "Berlin",
		},
	})
}

func (s *Suite) TestGetStructInfo(c *check.C) {
	result := GetStructInfo(people[0])

	c.Check(result, check.DeepEquals, map[string]FieldInfo{
		"name":         FieldInfo{Name: "Name", Unique: true},
		"address.city": FieldInfo{Name: "City", Index: true},
	})
}
