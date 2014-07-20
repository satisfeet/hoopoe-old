package validation

import (
	"fmt"

	"gopkg.in/validator.v1"
)

type Error struct {
	errors map[string][]error
}

func (err Error) Error() string {
	for k, errs := range err.errors {
		return fmt.Sprintf("%s value had %s", k, errs[0].Error())
	}

	return "no error"
}

func init() {
	validator.SetValidationFunc("min", Min)
	validator.SetValidationFunc("mail", Email)
}

func Validate(v interface{}) error {
	if ok, errs := validator.Validate(v); !ok {

		for _, errs := range errs {
			for _, err := range errs {
				switch err {
				case validator.ErrUnknownTag:
					return err
				case validator.ErrUnsupported:
					return err
				case validator.ErrBadParameter:
					return err
				}
			}
		}

		return Error{errs}
	}

	return nil
}
