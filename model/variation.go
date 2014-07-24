package model

import "github.com/satisfeet/hoopoe/model/validation"

// Variation represents the properties of a Product. We use it as there may be
// different derivates of a product commonly in color and size. It can further
// also store stock information as a variation is the smallest unit in stock.
type Variation struct {
	Size  string `json:"size" validate:"required,len=5"`
	Color string `json:"color" validate:"required,min=3"`
}

func (v Variation) Validate() error {
	return validation.Validate(v)
}
