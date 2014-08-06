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

type person struct {
	Name    string
	Address address
}

type address struct {
	City string
}

func (s *Suite) TestGetFieldValues(c *check.C) {
	p1 := person{"Bodo", address{"Berlin"}}
	p2 := person{}

	c.Check(GetFieldValues(p1), check.DeepEquals, map[string]interface{}{
		"name": "Bodo",
		"address": map[string]interface{}{
			"city": "Berlin",
		},
	})
	c.Check(GetFieldValues(p2), check.DeepEquals, map[string]interface{}{})
}
