package model

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/util"
)

type Product struct {
	Id          bson.ObjectId   `bson:"_id"`
	Name        string          `validate:"required,min=10,max=20"`
	Images      []bson.ObjectId `validate:"min=1"`
	Pricing     Pricing         `validate:"required,nested"`
	Variations  []Variation     `validate:"required,nested"`
	Description string          `validate:"required,min=40"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(util.GetFieldValues(p))
}

type Variation struct {
	Size  string `validate:"required,len=5"`
	Color string `validate:"required,min=3"`
}
