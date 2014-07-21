package validation

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/validator.v1"
)

func TestValidation(t *testing.T) {
	check.Suite(&ValidationSuite{})
	check.TestingT(t)
}

type Value struct {
	Foo interface{} `validate:"min=1"`
}

type ValidationSuite struct{}

func (s *ValidationSuite) TestValidate(c *check.C) {
	c.Check(Validate(Value{""}), check.FitsTypeOf, Error{})
	c.Check(Validate(Value{"bar"}), check.IsNil)
	c.Check(Validate(Value{[]int{}}), check.Equals, validator.ErrUnsupported)
}
