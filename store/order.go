package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
)

type Order struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Items    []OrderItem   `json:"items" validate:"nested"`
	Pricing  Pricing       `json:"pricing" validate:"nested"`
	Customer mgo.DBRef     `json:"customer" validate:"required"`
}

func (o Order) Validate() error {
	return validation.Validate(o)
}

type OrderItem struct {
	Product   mgo.DBRef `json:"product" validate:"required,ref"`
	Quantity  int       `json:"quantity" validate:"required"`
	Pricing   Pricing   `json:"price" validate:"nested"`
	Variation Variation `json:"variation" validate:"nested"`
}

func (i OrderItem) Validate() error {
	return validation.Validate(i)
}
