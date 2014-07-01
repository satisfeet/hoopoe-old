package customers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	collection *mgo.Collection
)

// Setup shared database instance and ensure collection index.
func Setup(db *mgo.Database) {
	collection = db.C("customers")

	collection.EnsureIndex(mgo.Index{
		Key: []string{
			"email",
		},
		Unique: true,
	})
	collection.EnsureIndex(mgo.Index{
		Key: []string{
			"name",
			"company",
			"address.city",
			"address.street",
		},
	})
}

// Alias to condition map.
type Query map[string]string

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

// Inserts a new Customer into collection.
func Create(customer *Customer) error {
	if !customer.Id.Valid() {
		customer.Id = bson.NewObjectId()
	}

	return collection.Insert(customer)
}

// Updates an existing Customer to collection.
func Update(customer *Customer) error {
	return collection.UpdateId(customer.Id, customer)
}

// Removes an existing Customer from collection.
func Remove(customer *Customer) error {
	return collection.RemoveId(customer.Id)
}

// Returns array of Customers matching conditions where "conditions"
// may have property "search" for pseudo text search accross index.
func FindAll(query Query) ([]Customer, error) {
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

	return c, collection.Find(q).All(&c)
}

// Returns instance of Customer matching conditions where "conditions"
// should have property "id" for equals ObjectIdHex query.
func FindOne(query Query) (Customer, error) {
	c := Customer{}

	q := bson.M{
		"_id": bson.ObjectIdHex(query["id"]),
	}

	return c, collection.Find(q).One(&c)
}
