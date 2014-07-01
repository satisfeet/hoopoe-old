package store

import (
	"labix.org/v2/mgo"

	"github.com/satisfeet/hoopoe/store/customers"
)

func Init(c map[string]string) error {
	s, err := mgo.Dial(c["mongo"])

	if err != nil {
		return err
	}

	customers.Open(s.DB(""))

	return nil
}
