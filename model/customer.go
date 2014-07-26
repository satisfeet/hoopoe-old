package model

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
)

// Fields which should be on the index. These fields can be safely used by for
// full-text-search on documents.
var CustomerIndex = []string{
	"name",
	"email",
	"company",
	"address.city",
	"address.street",
}

// Customer represents ... a customer.
type Customer struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name,omitempty" validate:"required,min=5,max=40"`
	Email   string        `json:"email,omitempty" validate:"required,email"`
	Company string        `json:"company,omitempty" validate:"min=5,max=40"`
	Address Address       `json:"address,omitempty" validate:"required,nested"`
}

// Returns errors if mandatory fields are missing or invalid.
//
// TODO: It MAY be a good idea to do an DNS check on the email address.
func (c Customer) Validate() error {
	return validation.Validate(c)
}
