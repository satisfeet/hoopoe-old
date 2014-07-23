package model

import (
	"github.com/satisfeet/hoopoe/model/validation"
	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id          bson.ObjectId      `json:"id"     bson:"_id"`
	Name        string             `json:"name"`
	Pricing     Pricing            `json:"pricing"`
	Variations  []ProductVariation `json:"variations"`
	Description string             `json:"description"`
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

type ProductVariation struct {
	Size  string `json:"size"`
	Color string `json:"color"`
}

func (p ProductVariation) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(p.Size); err != nil {
		errs.Set("size", err)
	}
	if err := validation.Required(p.Color); err != nil {
		errs.Set("color", err)
	}

	if errs.Has() {
		return errs
	}
	return nil
}
