package model

import "github.com/satisfeet/hoopoe/model/validation"

type Address struct {
	City   string `json:"city"   "omitempty" validate:"nonzero,min=4,max=40"`
	Street string `json:"street" "omitempty" validate:"min=6,max=60"`
	Zip    int    `json:"zip"    "omitempty" validate:"min=10000, max=99999"`
}

func (a Address) Validate() error {
	return validation.Validate(a)
}
