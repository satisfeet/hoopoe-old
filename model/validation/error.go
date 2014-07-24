package validation

import "fmt"

type Error map[string][]error

func (err Error) Error() string {
	for k, errs := range err {
		return fmt.Sprintf("%s has %s", k, Errors(errs))
	}
	return ""
}

type Errors []error

func (errs Errors) Add(err error) Errors {
	return append(errs, err)
}

func (errs Errors) Error() string {
	if len(errs) > 0 {
		return errs[0].Error()
	}
	return ""
}
