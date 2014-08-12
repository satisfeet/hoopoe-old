package mongo

import (
	"io"

	"gopkg.in/mgo.v2/bson"
)

type File struct {
	Id bson.ObjectId

	io.ReadWriteCloser
}

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
