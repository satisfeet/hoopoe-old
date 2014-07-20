package validation

import (
	"testing"

	"gopkg.in/validator.v1"
)

func TestMin(t *testing.T) {
	if v := Min(5, "4"); v != nil {
		t.Errorf("Expected error to be nil but it was %s.\n", v)
	}
	if v := Min("abc", "2"); v != nil {
		t.Errorf("Expected error to be nil but it was %s.\n", v)
	}
	if v := Min(1, "ab"); v != validator.ErrBadParameter {
		t.Errorf("Expected error to be ErrBadParam but it was %s.\n", v)
	}
	if v := Min(true, "12"); v != validator.ErrUnsupported {
		t.Errorf("Expected error to be ErrBadValue but it was %s.\n", v)
	}
	if v := Min(10, "11"); v != validator.ErrMin {
		t.Errorf("Expected error to be ErrTooSmall but it was %s.\n", v)
	}
	if v := Min("abc", "4"); v != validator.ErrMin {
		t.Errorf("Expected error to be ErrTooSmall but it was %s.\n", v)
	}
}

func TestEmail(t *testing.T) {
	if v := Email("i@foo.me", ""); v != nil {
		t.Errorf("Expected error to be nil but it was %s.\n", v)
	}
	if v := Email(1234, ""); v != validator.ErrUnsupported {
		t.Errorf("Expected error to be ErrBadValue but it was %s.\n", v)
	}
	if v := Email("@foobar.me", ""); v != validator.ErrInvalid {
		t.Errorf("Expected error to be ErrInvalidEmail but it was %s.\n", v)
	}
}
