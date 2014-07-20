package validation

import (
	"net/mail"
	"reflect"
	"strconv"

	"gopkg.in/validator.v1"
)

func Min(v interface{}, s string) error {
	var valid bool

	switch r := reflect.ValueOf(v); r.Kind() {
	case reflect.String:
		i, err := parseInt(s)

		if err != nil {
			return validator.ErrBadParameter
		}

		valid = len(r.String()) > int(i)
	case reflect.Int:
		i, err := parseInt(s)

		if err != nil {
			return validator.ErrBadParameter
		}

		valid = r.Int() > int64(i)
	case reflect.Float64:
		i, err := parseFloat(s)

		if err != nil {
			return validator.ErrBadParameter
		}

		valid = r.Float() > i
	default:
		return validator.ErrUnsupported
	}

	if !valid {
		return validator.ErrMin
	}
	return nil
}

func Email(v interface{}, _ string) error {
	var valid bool

	switch r := reflect.ValueOf(v); r.Kind() {
	case reflect.String:
		_, err := mail.ParseAddress(r.String())

		valid = err == nil
	default:
		return validator.ErrUnsupported
	}

	if !valid {
		return validator.ErrInvalid
	}
	return nil
}

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 64)
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
