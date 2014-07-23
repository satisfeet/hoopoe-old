package model

import "gopkg.in/mgo.v2/bson"

var (
	// Fields which are on the index and searchable.
	CustomerIndex = []string{
		"name",
		"email",
		"company",
		"address.city",
		"address.street",
	}
)

type Customer struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name"                validate:"nonzero,min=5,max=40"`
	Email   string        `json:"email"               validate:"nonzero,email"`
	Company string        `json:"company" "omitempty" validate:"min=6,max=50"`
	Address Address       `json:"address" "omitempty"`
}
