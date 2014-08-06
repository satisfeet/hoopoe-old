package model

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestAddress(t *testing.T) {
	check.Suite(&AddressSuite{})
	check.TestingT(t)
}

type AddressSuite struct{}

func (s *AddressSuite) TestMarshal(c *check.C) {
	c.Check(Address{}.Marshal(), check.DeepEquals, map[string]interface{}{})

	c.Check(Address{
		City:    "Berlin",
		Street:  "Bundesallee 22",
		ZipCode: 10999,
	}.Marshal(), check.DeepEquals, map[string]interface{}{
		"city":    "Berlin",
		"street":  "Bundesallee 22",
		"zipcode": 10999,
	})
}
