package model

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          bson.ObjectId   `bson:"_id"`
	Name        string          `validate:"required,min=10,max=20"`
	Images      []bson.ObjectId `validate:"required,min=1"`
	Pricing     Pricing         `validate:"required,nested"`
	Variations  []Variation     `validate:"required,nested"`
	Description string          `validate:"required,min=60"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type Variation struct {
	Size  string `validate:"required,len=5"`
	Color string `validate:"required,min=3"`
}
