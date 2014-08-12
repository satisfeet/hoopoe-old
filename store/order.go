package store

import (
	"fmt"

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

type Pricing struct {
	Retail int64 `json:"retail" validate:"required,min=1"`
}

func (p Pricing) Float() float64 {
	return float64(p.Retail) / 100
}

func (p Pricing) String() string {
	return fmt.Sprintf("%.2f €", p.Float())
}
