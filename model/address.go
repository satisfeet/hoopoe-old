package model

type Address struct {
	City    string `validate:"required,min=4,max=60" store:"index"`
	Street  string `validate:"min=4,max=60" store:"index"`
	ZipCode int    `validate:"min=10000,max=99999"`
}
