package model

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestProduct(t *testing.T) {
	check.Suite(&ProductSuite{})
	check.TestingT(t)
}

type ProductSuite struct{}

func (s *ProductSuite) TestValidate(c *check.C) {
	p := Pricing{
		Retail: 299,
	}
	v := []Variation{
		Variation{
			Color: "black",
			Size:  "42-45",
		},
	}

	c.Check(Product{
		Name:       "Business Socks",
		Pricing:    p,
		Variations: v,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.IsNil)

	c.Check(Product{
		Pricing:    p,
		Variations: v,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.ErrorMatches, "Name.*")

	c.Check(Product{
		Name:       "Business Socks",
		Variations: v,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.ErrorMatches, "Pricing.*")

	c.Check(Product{
		Name:    "Business Socks",
		Pricing: p,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.ErrorMatches, "Variations.*")

	c.Check(Product{
		Name:       "Business Socks",
		Pricing:    p,
		Variations: v,
	}.Validate(), check.ErrorMatches, "Description.*")
}
