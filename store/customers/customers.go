package customers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	db *mgo.Database
)

// Setup shared database instance and ensure collection indices.
func Setup(session *mgo.Session) {
	db = session.DB("")

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

// Represents a Customer specific database query.
type Query bson.M

// Appends an equals hex ObjectId condition to query.
func (q *Query) Id(param string) {
	if len(param) != 0 {
		(*q)["_id"] = bson.ObjectIdHex(param)
	}
}

// Appends a pseudo text search query to indexed fields.
func (q *Query) Search(param string) {
	searchable := []string{
		"name",
		"email",
		"company",
		"address.city",
		"address.street",
	}

	if len(param) != 0 {
		o := []bson.M{}
		r := bson.RegEx{param, "i"}

		for _, value := range searchable {
			c := bson.M{}
			c[value] = &r
			o = append(o, c)
		}

		(*q)["$or"] = o
	}
}

// Inserts a new Customer into collection.
func Create(customer *Customer) error {
	if !customer.Id.Valid() {
		customer.Id = bson.NewObjectId()
	}

	return db.C("customers").Insert(customer)
}

// Updates an existing Customer to collection.
func Update(customer *Customer) error {
	return db.C("customers").UpdateId(customer.Id, customer)
}

// Removes an existing Customer from collection.
func Remove(customer *Customer) error {
	return db.C("customers").RemoveId(customer.Id)
}

// Returns array of Customers matching Query conditions.
func FindAll(query *Query) ([]Customer, error) {
	r := []Customer{}

	return r, db.C("customers").Find(query).All(&r)
}

// Returns instance of Customer matching Query conditions.
func FindOne(query *Query) (Customer, error) {
	r := Customer{}

	return r, db.C("customers").Find(query).One(&r)
}
