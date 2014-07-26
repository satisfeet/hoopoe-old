package store

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/store/mongodb"
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

type CustomerStore struct {
	Mongo common.Store
}

func NewCustomerStore() *CustomerStore {
	return &CustomerStore{
		Mongo: mongodb.DefaultStore,
	}
}

func (s *CustomerStore) All(m *[]Customer) error {
	q := mongodb.Query{}

	return s.Mongo.Find(q, m)
}

func (s *CustomerStore) One(i string, m *Customer) error {
	q := mongodb.Query{}
	if err := q.Id(i); err != nil {
		return err
	}

	return s.Mongo.FindOne(q, m)
}

func (s *CustomerStore) Search(m *[]Customer) error {
	q := mongodb.Query{}

	return s.Mongo.Find(q, m)
}

func (s *CustomerStore) Insert(c *Customer) error {
	if !c.Id.Valid() {
		c.Id = bson.NewObjectId()
	}
	if err := c.Validate(); err != nil {
		return err
	}

	return s.Mongo.Insert(c)
}

func (s *CustomerStore) Update(c *Customer) error {
	if err := c.Validate(); err != nil {
		return err
	}

	return s.Mongo.Update(c)
}

func (s *CustomerStore) Remove(c *Customer) error {
	return s.Mongo.Remove(c)
}
