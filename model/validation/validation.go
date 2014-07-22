package validation

import (
	"net/mail"
	"reflect"
	"strings"
)

type Error string

func (err Error) Error() string {
	return string(err)
}

var (
	ErrEmail    = Error("invalid email")
	ErrRange    = Error("invalid range")
	ErrLength   = Error("invalid length")
	ErrRequired = Error("required")
)

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
	if v == reflect.Zero(reflect.TypeOf(v)).Interface() {
		return ErrRequired
	}
	return nil
}
