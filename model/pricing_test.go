package model

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestPricing(t *testing.T) {
	check.Suite(&PricingSuite{})
	check.TestingT(t)
}

type PricingSuite struct{}

func (s *PricingSuite) TestValidate(c *check.C) {
	c.Check(Pricing{
		Retail: 001,
	}.Validate(), check.IsNil)

	c.Check(Pricing{
		Retail: -11,
	}.Validate(), check.ErrorMatches, "retail has invalid range")

	c.Check(Pricing{}.Validate(), check.ErrorMatches, "retail has invalid value")
}
