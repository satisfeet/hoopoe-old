package model

type Validatable interface {
	Validate() error
}
