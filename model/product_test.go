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
	v := []ProductVariation{
		ProductVariation{
			Color: "black",
			Size:  "42",
		},
	}

	c.Check(Product{
		Name:        "Business Socks",
		Pricing:     p,
		Variations:  v,
		Description: "These are some Business socks!",
	}.Validate(), check.IsNil)

	c.Check(Product{
		Pricing:     p,
		Variations:  v,
		Description: "These are some Business socks!",
	}.Validate(), check.ErrorMatches, "name has invalid.*")

	c.Check(Product{
		Name:        "Business Socks",
		Variations:  v,
		Description: "These are some Business socks!",
	}.Validate(), check.ErrorMatches, "pricing has .*")

	c.Check(Product{
		Name:        "Business Socks",
		Pricing:     p,
		Description: "These are some Business socks!",
	}.Validate(), check.ErrorMatches, "variations has .*")

	c.Check(Product{
		Name:       "Business Socks",
		Pricing:    p,
		Variations: v,
	}.Validate(), check.ErrorMatches, "description has invalid.*")
}

func TestProductVariation(t *testing.T) {
	check.Suite(&ProductVariationSuite{})
	check.TestingT(t)
}

type ProductVariationSuite struct{}

func (s *ProductVariationSuite) TestValidate(c *check.C) {
	c.Check(ProductVariation{
		Color: "red",
		Size:  "42",
	}.Validate(), check.IsNil)

	c.Check(ProductVariation{
		Color: "red",
	}.Validate(), check.ErrorMatches, "size has invalid .*")

	c.Check(ProductVariation{
		Size: "42",
	}.Validate(), check.ErrorMatches, "color has invalid .*")
}
