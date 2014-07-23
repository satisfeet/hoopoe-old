package model

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestCustomer(t *testing.T) {
	check.Suite(&CustomerSuite{})
	check.TestingT(t)
}

type CustomerSuite struct{}

func (s *CustomerSuite) TestValidate(c *check.C) {
	a := Address{
		City: "Berlin",
	}

	c.Check(Customer{
		Name:    "Bodo Kaiser",
		Email:   "i@bodokaiser.io",
		Address: a,
	}.Validate(), check.IsNil)
	c.Check(Customer{
		Name:    "Bodo Kaiser",
		Email:   "i@bodokaiser.io",
		Address: a,
		Company: "satisfeet",
	}.Validate(), check.IsNil)

	c.Check(Customer{
		Email:   "foo@bar.org",
		Address: a,
	}.Validate(), check.ErrorMatches, "name has invalid.*")
	c.Check(Customer{
		Name:    "Bodo Kaiser",
		Address: a,
	}.Validate(), check.ErrorMatches, "email has invalid.*")
	c.Check(Customer{
		Name:  "Bodo Kaiser",
		Email: "foo@bar.org",
	}.Validate(), check.ErrorMatches, "address has invalid.*")
}
