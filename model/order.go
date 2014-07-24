package model

import (
	"github.com/satisfeet/hoopoe/model/validation"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Items    []OrderItem   `json:"items" validate:"required"`
	Pricing  Pricing       `json:"pricing"`
	Customer mgo.DBRef     `json:"customer" validate:"required"`
}

func (o Order) Validate() error {
	if err := validation.Validate(o); err != nil {
		return err
	}
	return o.Pricing.Validate()
}

func (o Order) SetCustomer(c Customer) {
	o.Customer = mgo.DBRef{
		Id:         c.Id,
		Collection: "customers",
	}
}

type OrderItem struct {
	Product   mgo.DBRef `json:"product" validate:"required"`
	Quantity  int       `json:"quantity" validate:"required,min=1"`
	Pricing   Pricing   `json:"price"`
	Variation Variation `json:"variation"`
}

func (i OrderItem) Validate() error {
	if err := i.Pricing.Validate(); err != nil {
		return err
	}
	if err := i.Variation.Validate(); err != nil {
		return err
	}
	return validation.Validate(i)
}

func (i OrderItem) SetProduct(p Product) {
	i.Product = mgo.DBRef{
		Id:         p.Id,
		Collection: "products",
	}
}
