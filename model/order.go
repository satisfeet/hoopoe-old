package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
)

type ErrInvalidOrder string

func (err ErrInvalidOrder) Error() string {
	return string(err)
}

type Order struct {
	Id         bson.ObjectId `json:"id" bson:"_id"`
	Items      []OrderItem   `json:"items" validate:"nested"`
	Pricing    Pricing       `json:"pricing" validate:"nested"`
	Customer   Customer      `json:"-" bson:"-"`
	CustomerId bson.ObjectId `json:"customer" validate:"required"`
	Number     int           `json:"number", validate:"required,min=1"`
}

func (o Order) Validate() error {
	if err := validation.Validate(o.Pricing); err != nil {
		return ErrInvalidOrder("Pricing " + err.Error())
	}
	if !o.CustomerId.Valid() {
		return ErrInvalidOrder("CustomerId required")
	}
	if o.Number <= 0 {
		return ErrInvalidOrder("Number must be bigger than 1")
	}

	return nil
}

type OrderItem struct {
	Product   Product       `json:"-" bson:"-"`
	ProductId bson.ObjectId `json:"product" validate:"required"`
	Quantity  int           `json:"quantity" validate:"required"`
	Pricing   Pricing       `json:"pricing" validate:"nested"`
	Variation Variation     `json:"variation" validate:"nested"`
}

func (oi OrderItem) Validate() error {
	if err := validation.Valid(oi.ProductId, "required"); err != nil {
		return ErrInvalidOrder("ProductId " + err.Error())
	}
	if err := validation.Valid(oi.Quantity, "required"); err != nil {
		return ErrInvalidOrder("Qantity " + err.Error())
	}
	if err := validation.Validate(oi.Pricing); err != nil {
		return err
	}
	if err := validation.Validate(oi.Variation); err != nil {
		return err
	}

	return nil
}
