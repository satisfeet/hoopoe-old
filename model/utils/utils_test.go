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
	Name    string
	Address Address
}

type Address struct {
	City string
}

func TestUtils(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct{}

func (s *Suite) TestGetFieldValues(c *check.C) {
	c.Check(GetFieldValues(people[0]), check.DeepEquals, map[string]interface{}{})
	c.Check(GetFieldValues(people[1]), check.DeepEquals, map[string]interface{}{
		"name": "Bodo",
		"address": map[string]interface{}{
			"city": "Berlin",
		},
	})
}

func (s *Suite) TestGetNestedFieldPointer(c *check.C) {
	p1 := GetNestedFieldPointer(&people[1], "Name")
	p2 := GetNestedFieldPointer(&people[1], "City")

	v1, ok1 := p1.(*string)
	v2, ok2 := p2.(*string)

	c.Assert(ok1, check.Equals, true)
	c.Assert(ok2, check.Equals, true)

	c.Check(v1, check.Equals, &people[1].Name)
	c.Check(v2, check.Equals, &people[1].Address.City)
}
