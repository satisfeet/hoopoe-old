package model

import "github.com/satisfeet/hoopoe/model/validation"

type Pricing struct {
	Retail int64 `json:"retail"`
}

func (p Pricing) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(p.Retail); err == nil {
		if err := validation.Range(int(p.Retail), 1, 0); err != nil {
			errs.Set("retail", err)
		}
	} else {
		errs.Set("retail", err)
	}

	if errs.Has() {
		return errs
	}
	return nil
}
