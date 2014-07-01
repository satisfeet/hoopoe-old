package store

import (
	"labix.org/v2/mgo"

	"github.com/satisfeet/hoopoe/store/customers"
)

// Opens a new connection session to mongodb and setups subpackages.
func Open(config map[string]string) error {
	s, err := mgo.Dial(config["mongo"])

	if err != nil {
		return err
	}

	d := s.DB("")

	customers.Setup(d)

	return err
}
