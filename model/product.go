package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model/validation"
)

type Product struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Name        string        `json:"name" validate:"required,min=10,max=20"`
	Pricing     Pricing       `json:"pricing"`
	Variations  []Variation   `json:"variations" validate:"required,min=1"`
	Description string        `json:"description" validate:"required,min=60"`
}

func (p Product) Validate() error {
	if err := validation.Validate(p); err != nil {
		return err
	}
	for _, v := range p.Variations {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return p.Pricing.Validate()
}
