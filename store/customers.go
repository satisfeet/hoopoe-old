package store

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// Represents a Customer instance.
type Customer struct {
	Id      bson.ObjectId   `bson:"_id"     json:"id"`
	Name    string          `bson:"name"    json:"name"`
	Email   string          `bson:"email"   json:"email"`
	Company string          `bson:"company" json:"company,omitempty"`
	Address CustomerAddress `bson:"address" json:"address"`
}

// Represents a Customer instance's Address struct.
type CustomerAddress struct {
	Street string `bson:"street"  json:"street,omitempty"`
	City   string `bson:"city"    json:"city,omitempty"`
	Zip    uint16 `bson:"zip"     json:"zip,omitempty"`
}

// Setup customers index.
func CustomersIndex() {
	db.C("customers").EnsureIndex(mgo.Index{
		Key: []string{
			"email",
		},
		Unique: true,
	})
	db.C("customers").EnsureIndex(mgo.Index{
		Key: []string{
			"name",
			"company",
			"address.city",
			"address.street",
		},
	})
}

// Inserts a new Customer into db.C("customers").
func CustomersCreate(customer *Customer) error {
	if !customer.Id.Valid() {
		customer.Id = bson.NewObjectId()
	}

	return db.C("customers").Insert(customer)
}

// Updates an existing Customer to db.C("customers").
func CustomersUpdate(customer *Customer) error {
	return db.C("customers").UpdateId(customer.Id, customer)
}

// Removes an existing Customer from db.C("customers").
func CustomersRemove(customer *Customer) error {
	return db.C("customers").RemoveId(customer.Id)
}

// Returns array of Customers matching conditions where "conditions"
// may have property "search" for pseudo text search accross index.
func CustomersFindAll(query Query) ([]Customer, error) {
	c := []Customer{}

	q := bson.M{}

	if len(query["search"]) != 0 {
		r := bson.RegEx{query["search"], "i"}

		q["$or"] = []bson.M{
			bson.M{"name": r},
			bson.M{"email": r},
			bson.M{"company": r},
			bson.M{"address.city": r},
			bson.M{"address.street": r},
		}
	}

	return c, db.C("customers").Find(q).All(&c)
}

// Returns instance of Customer matching conditions where "conditions"
// should have property "id" for equals ObjectIdHex query.
func CustomersFindOne(query Query) (Customer, error) {
	c := Customer{}

	q := bson.M{
		"_id": bson.ObjectIdHex(query["id"]),
	}

	return c, db.C("customers").Find(q).One(&c)
}
