package model

import (
	"strconv"

	"github.com/satisfeet/hoopoe/model/validation"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	Id       bson.ObjectId `json:"id"     bson:"_id"`
	Items    []OrderItem   `json:"items"`
	Pricing  Pricing       `json:"pricing"`
	Customer mgo.DBRef     `json:"customer"`
}

func (o Order) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(o.Customer.Id); err != nil {
		errs.Set("customer", err)
	}
	if err := validation.Required(o.Items); err != nil {
		errs.Set("items", err)
	}
	if err := o.Pricing.Validate(); err != nil {
		errs.Set("pricing", err)
	}

	for n, i := range o.Items {
		if err := i.Validate(); err != nil {
			errs.Set("item #"+strconv.Itoa(n), err)

			break
		}
	}

	if errs.Has() {
		return errs
	}
	return nil
}

func (o Order) SetCustomer(c Customer) {
	o.Customer = mgo.DBRef{
		Id:         c.Id,
		Collection: "customers",
	}
}

type OrderItem struct {
	Product   mgo.DBRef `json:"product"`
	Quantity  int       `json:"quantity"`
	Pricing   Pricing   `json:"price"`
	Variation Variation `json:"variation"`
}

func (i OrderItem) Validate() error {
	errs := validation.Error{}

	if err := validation.Required(i.Product.Id); err != nil {
		errs.Set("product", err)
	}
	if err := validation.Required(i.Quantity); err == nil {
		if err := validation.Range(i.Quantity, 1, 0); err != nil {
			errs.Set("quantity", err)
		}
	} else {
		errs.Set("quantity", err)
	}
	if err := i.Pricing.Validate(); err != nil {
		errs.Set("pricing", err)
	}
	if err := i.Variation.Validate(); err != nil {
		errs.Set("variation", err)
	}

	if errs.Has() {
		return errs
	}
	return nil
}

func (i OrderItem) SetProduct(p Product) {
	i.Product = mgo.DBRef{
		Id:         p.Id,
		Collection: "products",
	}
}
