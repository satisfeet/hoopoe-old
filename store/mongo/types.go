package mongo

import (
	"io"

	"gopkg.in/mgo.v2/bson"
)

type File struct {
	Id bson.ObjectId

	io.ReadWriteCloser
}

func IdFromString(id string) bson.ObjectId {
	var oid bson.ObjectId

	if bson.IsObjectIdHex(id) {
		oid = bson.ObjectIdHex(id)
	}

	return oid
}
