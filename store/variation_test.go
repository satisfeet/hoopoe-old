package store

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestVariation(t *testing.T) {
	check.Suite(&VariationSuite{})
	check.TestingT(t)
}

type VariationSuite struct{}

func (s *VariationSuite) TestValidate(c *check.C) {
	c.Check(Variation{
		Color: "red",
		Size:  "42-44",
	}.Validate(), check.IsNil)

	c.Check(Variation{
		Color: "red",
	}.Validate(), check.ErrorMatches, "Size.*")

	c.Check(Variation{
		Size: "42-44",
	}.Validate(), check.ErrorMatches, "Color.*")
}
