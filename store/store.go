package store

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

var ErrBadParam = errors.New("bad param")

func ParseId(id interface{}) bson.ObjectId {
	var oid bson.ObjectId

	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			oid = bson.ObjectIdHex(t)
		}
	case bson.ObjectId:
		if t.Valid() {
			oid = t
		}
	}

	return oid
}
