package store

import "labix.org/v2/mgo/bson"

type Query bson.M

func (q Query) Id(param string) {
	q["_id"] = bson.ObjectIdHex(param)
}

func (q Query) Search(param string) {
	r := bson.RegEx{param, "i"}

	q["$or"] = []bson.M{
		{"name": r},
		{"email": r},
		{"company": r},
		{"address.city": r},
		{"address.street": r},
	}
}
