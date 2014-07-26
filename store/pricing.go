package store

import "github.com/satisfeet/go-validation"

type Pricing struct {
	Retail int64 `json:"retail" validate:"required,min=1"`
}

func (p Pricing) Validate() error {
	return validation.Validate(p)
}
