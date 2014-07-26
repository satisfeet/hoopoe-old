package store

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
	i := []string{
		"24781374dhsfhkhk",
	}
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
		Images:     i,
		Pricing:    p,
		Variations: v,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.IsNil)

	c.Check(Product{
		Images:     i,
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
		Pricing:    p,
		Variations: v,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.ErrorMatches, "Images.*")

	c.Check(Product{
		Name:       "Business Socks",
		Images:     i,
		Variations: v,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.ErrorMatches, "Pricing.*")

	c.Check(Product{
		Name:    "Business Socks",
		Images:  i,
		Pricing: p,
		Description: `
			These are some Business socks!
			They are really really nice.
			You should try them!
		`,
	}.Validate(), check.ErrorMatches, "Variations.*")

	c.Check(Product{
		Name:       "Business Socks",
		Images:     i,
		Pricing:    p,
		Variations: v,
	}.Validate(), check.ErrorMatches, "Description.*")
}
