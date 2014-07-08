package store

import (
	"errors"

	"labix.org/v2/mgo/bson"
)

type Query bson.M

func (q Query) Id(hex string) error {
	if !bson.IsObjectIdHex(hex) {
		return errors.New("Invalid ObjectId.")
	}

	q["_id"] = bson.ObjectIdHex(hex)

	return nil
}

func (q Query) Search(param string, fields []string) {
	r := bson.RegEx{param, "i"}

	o := make([]bson.M, len(fields))

	for i := 0; i < len(fields); i++ {
		o[i] = make(bson.M, 1)
		o[i][fields[i]] = r
	}

	q["$or"] = o
}
