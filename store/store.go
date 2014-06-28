package store

import (
	"labix.org/v2/mgo"

	"github.com/satisfeet/hoopoe/store/customers"
)

func Open(config map[string]string) (err error) {
	s, err := mgo.Dial(config["mongo"])

	if err != nil {
		return
	}

	customers.Setup(s)

	return
}
