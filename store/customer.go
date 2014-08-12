package store

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/mongo"
	"github.com/satisfeet/hoopoe/utils"
)

type Address struct {
	City    string `validate:"required,min=4,max=60" store:"index"`
	Street  string `validate:"min=4,max=60" store:"index"`
	ZipCode int    `validate:"min=10000,max=99999"`
}

type Customer struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `validate:"required,min=5" store:"unique"`
	Email   string        `validate:"required,email" store:"unique"`
	Company string        `validate:"min=5,max=40" store:"index"`
	Address Address       `validate:"required,nested"`
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(c))
}

type CustomerStore struct {
	*store
}

func NewCustomerStore(s *mongo.Store) *CustomerStore {
	return &CustomerStore{
		store: &store{s},
	}
}

func (s *CustomerStore) Search(keyword string, m *[]Customer) error {
	q := mongo.Query{}

	if len(keyword) > 0 {
		or := []bson.M{}

		for k, _ := range utils.GetStructInfo(m) {
			m := bson.M{}
			m[k] = bson.RegEx{keyword, "i"}

			or = append(or, m)
		}

		q["$or"] = or
	}

	return s.mongo.Find(getName(m), q, m)
}

func (s *CustomerStore) RemoveId(id interface{}) error {
	return s.mongo.RemoveId("customers", id)
}
