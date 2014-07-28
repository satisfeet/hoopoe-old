package store

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Customer struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name,omitempty" validate:"required,min=5,max=40"`
	Email   string        `json:"email,omitempty" validate:"required,email"`
	Company string        `json:"company,omitempty" validate:"min=5,max=40"`
	Address Address       `json:"address,omitempty" validate:"required,nested"`
}

func (c Customer) Validate() error {
	return validation.Validate(c)
}

// Customers store.
type Customers struct {
	Mongo *mongo.Store
}

var (
	// Collection name to use for databases.
	CustomersName = "customers"

	// Indexed fields to search.
	CustomersIndex = []string{
		"name",
		"email",
		"company",
		"address.city",
		"address.street",
	}
)

func (s *Customers) All(m *[]Customer) error {
	q := mongo.Query{}

	return s.Mongo.FindAll(CustomersName, q, m)
}

func (s *Customers) One(i string, m *Customer) error {
	q := mongo.Query{}

	if err := q.Id(i); err != nil {
		return err
	}

	return s.Mongo.FindOne(CustomersName, q, m)
}

func (s *Customers) Search(v string, m *[]Customer) error {
	q := mongo.Query{}

	for _, i := range CustomersIndex {
		qr := mongo.Query{}
		qr.Regex(i, v)

		q.Or(qr)
	}

	return s.Mongo.FindAll(CustomersName, q, m)
}

func (s *Customers) Insert(c *Customer) error {
	if !c.Id.Valid() {
		c.Id = bson.NewObjectId()
	}
	if err := c.Validate(); err != nil {
		return err
	}

	return s.Mongo.Insert(CustomersName, c)
}

func (s *Customers) Update(c *Customer) error {
	q := mongo.Query{}

	if err := q.Id(c.Id); err != nil {
		return err
	}
	if err := c.Validate(); err != nil {
		return err
	}

	return s.Mongo.Update(CustomersName, q, c)
}

func (s *Customers) Remove(c *Customer) error {
	q := mongo.Query{}

	if err := q.Id(c.Id); err != nil {
		return err
	}

	return s.Mongo.Remove(CustomersName, q)
}
