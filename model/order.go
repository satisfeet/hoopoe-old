package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Items       []OrderItem   `json:"items" validate:"nested"`
	Pricing     Pricing       `json:"pricing" validate:"nested"`
	Customer    Customer      `json:"-" bson:"-"`
	CustomerRef *mgo.DBRef    `json:"customer" validate:"required"`
}

type OrderItem struct {
	Product    Product    `json:"-" bson:"-"`
	ProductRef *mgo.DBRef `json:"product" validate:"required"`
	Quantity   int        `json:"quantity" validate:"required"`
	Pricing    Pricing    `json:"price" validate:"nested"`
	Variation  Variation  `json:"variation" validate:"nested"`
}
