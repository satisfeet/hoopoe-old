package model

import (
	"encoding/json"

	"github.com/satisfeet/hoopoe/utils"
)

type Address struct {
	City    string `validate:"required,min=4,max=60" store:"index"`
	Street  string `validate:"min=4,max=60" store:"index"`
	ZipCode int    `validate:"min=10000,max=99999"`
}

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(a))
}
