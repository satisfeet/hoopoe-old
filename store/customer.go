package store

import "gopkg.in/mgo.v2/bson"

var (
	CustomerDatabase   = ""
	CustomerCollection = "customers"

	CustomerUnique = []string{
		"email",
	}
	CustomerIndices = []string{
		"name",
		"company",
		"address.street",
		"address.city",
	}
)

type Customer struct {
	Id      bson.ObjectId   `json:"id"     bson:"_id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Company string          `json:"company,omitempty"`
	Address CustomerAddress `json:"address"`
}

type CustomerAddress struct {
	Zip    int    `json:"zip,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
}
