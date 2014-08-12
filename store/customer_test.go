package store

import "gopkg.in/check.v1"

var customers = []Customer{
	Customer{
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
		Address: Address{
			City: "Berlin",
		},
	},
	Customer{
		Name:  "Denzel Washington",
		Email: "denzel@example.com",
		Address: Address{
			City: "Hollywood",
		},
	},
}

func (s *Suite) TestCustomerSearch(c *check.C) {
	m := []Customer{}

	err := s.customer.Search("denzel", &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, customers[1:])
}

func (s *Suite) TestCustomerRemoveId(c *check.C) {
	err := s.customer.RemoveId(customers[0].Id)
	c.Assert(err, check.IsNil)
}
