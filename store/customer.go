package store

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	*store
}

func NewCustomer(db *mgo.Database) *Customer {
	return &Customer{
		store: &store{db},
	}
}

func (s *Customer) Search(keyword string, m Model) error {
	q := query{}

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
