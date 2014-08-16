package model

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"

	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          bson.ObjectId   `bson:"_id"`
	Name        string          `validate:"required,min=10,max=20"`
	Images      []bson.ObjectId `validate:"min=1"`
	Pricing     Pricing         `validate:"required,nested"`
	Variations  []Variation     `validate:"required,nested"`
	Description string          `validate:"required,min=40"`
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type Variation struct {
	Size  string `json:"size" validate:"required,len=5"`
	Color string `json:"color" validate:"required,min=3"`
}
