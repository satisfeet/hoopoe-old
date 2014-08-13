package store

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	*store
}

func NewCustomer(s *mongo.Store) *Customer {
	return &Customer{
		store: &store{s},
	}
}

func (s *Customer) Search(keyword string, m *[]model.Customer) error {
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
