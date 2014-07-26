package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
)

// Product represents a product in stock. It contains a few details to be
// presented to the customer.
type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" validate:"required,min=10,max=20"`
	Images      []string      `json:"image" validate:"required,min=1"`
	Pricing     Pricing       `json:"pricing" validate:"required,nested"`
	Variations  []Variation   `json:"variations" validate:"required,nested"`
	Description string        `json:"description" validate:"required,min=60"`
}

// Returns error if product is invalid.
//
// NOTE: As products will only be defined internal we can be more loose about
// validation in comparison to other entities.
func (p Product) Validate() error {
	return validation.Validate(p)
}
