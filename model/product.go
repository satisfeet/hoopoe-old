package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model/validation"
)

type Product struct {
	Id          bson.ObjectId `json:"id"     bson:"_id"`
	Name        string        `json:"name"`
	Pricing     Pricing       `json:"pricing"`
	Variations  []Variation   `json:"variations"`
	Description string        `json:"description"`
}

func (p Product) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(p.Name); err != nil {
		errs.Set("name", err)
	}
	if err := validation.Required(p.Description); err != nil {
		errs.Set("description", err)
	}
	if err := validation.Required(p.Variations); err != nil {
		errs.Set("variations", err)
	}
	if err := p.Pricing.Validate(); err != nil {
		errs.Set("pricing", err)
	}

	for _, v := range p.Variations {
		if err := v.Validate(); err != nil {
			errs.Set("variations", err)

			break
		}
	}

	if errs.Has() {
		return errs
	}
	return nil
}
