package model

import "github.com/satisfeet/hoopoe/model/validation"

type Variation struct {
	Size  string `json:"size"`
	Color string `json:"color"`
}

func (p Variation) Validate() error {
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
