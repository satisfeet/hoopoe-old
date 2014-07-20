package validation

import "testing"

type Value struct {
	Foo interface{} `validate:"min=1"`
}

func TestValidate(t *testing.T) {
	if v := Validate(Value{Foo: ""}); v == nil {
		t.Errorf("Expected error to be not nil but it was %s.\n", v)
	}
	if v := Validate(Value{Foo: "bar"}); v != nil {
		t.Errorf("Expected error to be nil but it was %s.\n", v)
	}
}
