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
	Name    string        `json:"name"`
	Email   string        `json:"email,omitempty"`
	Company string        `json:"company,omitempty"`
	Address Address       `json:"address,omitempty"`
}

func (c Customer) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(c.Name); err == nil {
		if err := validation.Length(c.Name, 5, 40); err != nil {
			errs.Set("name", err)
		}
	} else {
		errs.Set("name", err)
	}
	if err := validation.Required(c.Email); err == nil {
		if err := validation.Email(c.Email); err != nil {
			errs.Set("email", err)
		}
	} else {
		errs.Set("email", err)
	}
	if err := validation.Required(c.Address); err == nil {
		if err := c.Address.Validate(); err != nil {
			errs.Set("address", err)
		}
	} else {
		errs.Set("address", err)
	}
	if err := validation.Required(c.Company); err == nil {
		if err := validation.Length(c.Company, 5, 40); err != nil {
			errs.Set("company", err)
		}
	}

	if errs.Has() {
		return errs
	}
	return nil
}
