package model

import "gopkg.in/mgo.v2/bson"

var (
	// Name of the collection.
	CustomerName = "customers"

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
	Id      bson.ObjectId `json:"id"     bson:"_id"`
	Name    string        `json:"name"`
	Email   string        `json:"email"`
	Company string        `json:"company  "omitempty"`
	Address Address       `json:"address" "omitempty"`
}
