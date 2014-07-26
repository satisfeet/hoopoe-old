package model

import "github.com/satisfeet/go-validation"

// Address represents a Germany-based address. For private sales the least
// possible value is to only contain the city property for online sales or
// fields are required which must be validated by an additional layer on the
// front-end and optionally in the back-end.
type Address struct {
	City    string `json:"city,omitempty" validate:"required,min=4,max=60"`
	Street  string `json:"street,omitempty" validate:"min=4,max=40`
	ZipCode int    `json:"zipcode,omitempty" validate:"min=10000,max=99999"`
}

// Returns validation errors.
//
// NOTE: It may be a good idea to lookup the address for existence.
func (a Address) Validate() error {
	return validation.Validate(a)
}
