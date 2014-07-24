package model

import (
	"github.com/satisfeet/hoopoe/model/validation"
	"gopkg.in/mgo.v2/bson"
)

var (
	// Fields which are on the index and searchable.
	CustomerIndex = []string{
		"name",
		"email",
		"company",
		"address.city",
		"address.street",
	}
)

type Customer struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name,omitempty" validate:"required,min=5,max=40"`
	Email   string        `json:"email,omitempty" validate:"required,email"`
	Company string        `json:"company,omitempty" validate:"min=5,max=40"`
	Address Address       `json:"address,omitempty" validate:"required,nested"`
}

func (c Customer) Validate() error {
	return validation.Validate(c)
}
