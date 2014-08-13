package store

import (
	"gopkg.in/check.v1"

	"github.com/satisfeet/hoopoe/model"
)

var customers = []model.Customer{
	model.Customer{
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
		Address: model.Address{
			City: "Berlin",
		},
	},
	model.Customer{
		Name:  "Denzel Washington",
		Email: "denzel@example.com",
		Address: model.Address{
			City: "Hollywood",
		},
	},
}

func (s *Suite) TestCustomerSearch(c *check.C) {
	m := []model.Customer{}

	err := s.customer.Search("denzel", &m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.DeepEquals, customers[1:])
}

func (s *Suite) TestCustomerRemoveId(c *check.C) {
	err := s.customer.RemoveId(customers[0].Id)
	c.Assert(err, check.IsNil)
}
