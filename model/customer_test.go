package model

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func TestCustomer(t *testing.T) {
	check.Suite(&CustomerSuite{})
	check.TestingT(t)
}

type CustomerSuite struct{}

func (s *CustomerSuite) TestMarshal(c *check.C) {
	c.Check(Customer{}.Marshal(), check.DeepEquals, map[string]interface{}{
		"id": "",
	})

	id := bson.NewObjectId()
	c.Check(Customer{
		Id:    id,
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
		Address: Address{
			City:    "Berlin",
			Street:  "Bundesallee 22",
			ZipCode: 10999,
		},
	}.Marshal(), check.DeepEquals, map[string]interface{}{
		"id":    id.Hex(),
		"name":  "Bodo Kaiser",
		"email": "i@bodokaiser.io",
		"address": map[string]interface{}{
			"city":    "Berlin",
			"street":  "Bundesallee 22",
			"zipcode": 10999,
		},
	})
}
