package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model/validation"
)

type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" validate:"required,min=10,max=20"`
	Images      []string      `json:"image" validate:"required,min=1"`
	Pricing     Pricing       `json:"pricing" validate:"required,nested"`
	Variations  []Variation   `json:"variations" validate:"required,nested"`
	Description string        `json:"description" validate:"required,min=60"`
}

func (p Product) Validate() error {
	return validation.Validate(p)
}
