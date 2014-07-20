package model

import "github.com/satisfeet/hoopoe/model/validation"

type Pricing struct {
	Retail int64 `json:"retail" validate:"nonzero,min=1"`
}

func (p Pricing) Validate() error {
	return validation.Validate(p)
}
