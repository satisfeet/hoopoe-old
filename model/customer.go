package model

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"

	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `validate:"required,min=5" store:"unique"`
	Email   string        `validate:"required,email" store:"unique"`
	Company string        `validate:"min=5,max=40" store:"index"`
	Address Address       `validate:"required,nested"`
}

func (c Customer) Validate() error {
	return validation.Validate(c)
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(c))
}
