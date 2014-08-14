package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	*store
}

func NewCustomer(s *mgo.Session) *Customer {
	return &Customer{
		store: &store{
			session:  s,
			database: s.DB(""),
		},
	}
}

func (s *Customer) Search(keyword string, m *[]model.Customer) error {
	q := bson.M{}

	if len(keyword) > 0 {
		or := []bson.M{}

		for k, _ := range utils.GetStructInfo(m) {
			m := bson.M{}
			m[k] = bson.RegEx{keyword, "i"}

			or = append(or, m)
		}

		q["$or"] = or
	}

	return s.collection(m).Find(q).All(m)
}
