package validation

import (
	"errors"
	"net/mail"
	"reflect"
	"strings"
)

var (
	ErrEmail    = errors.New("invalid email")
	ErrRange    = errors.New("invalid range")
	ErrLength   = errors.New("invalid length")
	ErrRequired = errors.New("invalid value")
)

type Validatable interface {
	Validate() error
}

func Email(s string) error {
	if _, err := mail.ParseAddress(s); err != nil {
		return ErrEmail
	} else {
		// the above sees "foo@bar." as valid so do
		// additional check if "." is not last char
		if strings.LastIndex(s, ".") == len(s)-1 {
			return ErrEmail
		}
	}
	return nil
}

func Range(i int, a, b int) error {
	if a != 0 {
		if i < a {
			return ErrRange
		}
	}
	if b != 0 {
		if i > b {
			return ErrRange
		}
	}
	return nil
}

func Length(s string, a, b int) error {
	l := len(s)

	if a != 0 {
		if l < a {
			return ErrLength
		}
	}
	if b != 0 {
		if l > b {
			return ErrLength
		}
	}
	return nil
}

func Required(v interface{}) error {
	if v == nil {
		return ErrRequired
	}

	switch r := reflect.ValueOf(v); r.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array:
		if r.Len() == 0 {
			return ErrRequired
		} else {
			return nil
		}
	}

	if v == reflect.Zero(reflect.TypeOf(v)).Interface() {
		return ErrRequired
	}
	return nil
}
