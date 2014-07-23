package model

import "github.com/satisfeet/hoopoe/model/validation"

type Address struct {
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
	Zip    int    `json:"zip,omitempty"`
}

func (a Address) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(a.Street); err == nil {
		if err := validation.Length(a.Street, 6, 60); err != nil {
			errs.Set("street", err)
		}
	}
	if err := validation.Required(a.City); err == nil {
		if err := validation.Length(a.City, 4, 40); err != nil {
			errs.Set("city", err)
		}
	} else {
		errs.Set("city", err)
	}
	if err := validation.Required(a.Zip); err == nil {
		if err := validation.Range(a.Zip, 10000, 99999); err != nil {
			errs.Set("zip", err)
		}
	}

	if errs.Has() {
		return errs
	}
	return nil
}
