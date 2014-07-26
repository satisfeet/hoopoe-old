package store

import "github.com/satisfeet/go-validation"

type Variation struct {
	Size  string `json:"size" validate:"required,len=5"`
	Color string `json:"color" validate:"required,min=3"`
}

func (v Variation) Validate() error {
	return validation.Validate(v)
}
