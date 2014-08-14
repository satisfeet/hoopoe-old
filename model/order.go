package model

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
)

type Order struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Items       []OrderItem   `json:"items" validate:"nested"`
	Pricing     Pricing       `json:"pricing" validate:"nested"`
	Customer    Customer      `json:"-" bson:"-"`
	CustomerRef *mgo.DBRef    `json:"customer" validate:"required"`
}

func (o Order) Validate() error {
	if err := validation.Validate(o.Pricing); err != nil {
		return errors.New("Pricing " + err.Error())
	}
	if o.CustomerRef == nil {
		return errors.New("CustomerRef required")
	}

	return nil
}

type OrderItem struct {
	Product    Product    `json:"-" bson:"-"`
	ProductRef *mgo.DBRef `json:"product" validate:"required"`
	Quantity   int        `json:"quantity" validate:"required"`
	Pricing    Pricing    `json:"price" validate:"nested"`
	Variation  Variation  `json:"variation" validate:"nested"`
}

func (oi OrderItem) Validate() error {
	if err := validation.Valid(oi.ProductRef, "required"); err != nil {
		return errors.New("ProductRef " + err.Error())
	}
	if err := validation.Valid(oi.Quantity, "required"); err != nil {
		return errors.New("Qantity " + err.Error())
	}
	if err := validation.Validate(oi.Pricing); err != nil {
		return err
	}
	if err := validation.Validate(oi.Variation); err != nil {
		return err
	}

	return nil
}
