package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
)

// Order represents an order made by a customer.
type Order struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Items    []OrderItem   `json:"items" validate:"nested"`
	Pricing  Pricing       `json:"pricing" validate:"nested"`
	Customer mgo.DBRef     `json:"customer" validate:"required"`
}

// Returns errors if order is invalid.
//
// TODO: We MUST check the order for duplicates and correct pricing also we must
// check if the customer exists.
func (o Order) Validate() error {
	return validation.Validate(o)
}

// Assigns the given customer to the order.
func (o Order) SetCustomer(c Customer) {
	o.Customer = mgo.DBRef{
		Id:         c.Id,
		Collection: "customers",
	}
}

// OrderItem represents a purchased item. It wraps a product reference with
// values as quantity total price and choosen variation.
type OrderItem struct {
	Product   mgo.DBRef `json:"product" validate:"required,ref"`
	Quantity  int       `json:"quantity" validate:"required"`
	Pricing   Pricing   `json:"price" validate:"nested"`
	Variation Variation `json:"variation" validate:"nested"`
}

// Validates a single order item.
func (i OrderItem) Validate() error {
	return validation.Validate(i)
}

// Assignes the given product to the item.
func (i OrderItem) SetProduct(p Product) {
	i.Product = mgo.DBRef{
		Id:         p.Id,
		Collection: "products",
	}
}
