package model

import (
	"testing"

	"github.com/satisfeet/hoopoe/model/validation"
	"gopkg.in/check.v1"
)

func TestAddress(t *testing.T) {
	check.Suite(&AddressSuite{})
	check.TestingT(t)
}

type AddressSuite struct{}

func (s *AddressSuite) TestValidate(c *check.C) {
	c.Check(Address{
		City: "Berlin",
	}.Validate(), check.IsNil)
	c.Check(Address{
		City:   "Berlin",
		Street: "Geiserichstr. 3",
		Zip:    12105,
	}.Validate(), check.IsNil)

	c.Check(Address{}.Validate(), check.Equals, validation.ErrRequired)
}
