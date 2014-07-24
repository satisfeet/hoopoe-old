package model

import "github.com/satisfeet/hoopoe/model/validation"

// Pricing represents pricing information. At the moment this is only the retail
// price however in future it will contain tax information and possible
// discounts.
type Pricing struct {
	Retail int64 `json:"retail" validate:"required,min=1"`
}

// Returns error if retail not set or negative.
func (p Pricing) Validate() error {
	return validation.Validate(p)
}
