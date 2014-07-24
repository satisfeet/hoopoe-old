package validation

import "gopkg.in/validator.v1"

var (
	// The default validator we extend and use.
	//
	// We use an extra instance over here to avoid
	// collisions with the one exported by validator.
	DefaultValidator = validator.NewValidator()
)

// Extends DefaultValidator to use our own constraints
// in favor to the built-ins.
func init() {
	DefaultValidator.SetValidationFunc("min", Minimum)
	DefaultValidator.SetValidationFunc("email", Email)
	DefaultValidator.SetValidationFunc("required", Required)
}

// Every type which implements a Validate method is
// Validatable.
//
// This will mostly be implemented by models which use
// one of the validation methods of this package.
type Validatable interface {
	Validate() error
}

// Returns Errors if validation of single value fails else
// it will return nil.
func Valid(v interface{}, s string) error {
	if ok, err := DefaultValidator.Valid(v, s); !ok {
		return Errors(err)
	}
	return nil
}

// Returns Error if validation of tagged struct value fails
// else it will return nil.
func Validate(v interface{}) error {
	if ok, err := DefaultValidator.Validate(v); !ok {
		return Error(err)
	}
	return nil
}
