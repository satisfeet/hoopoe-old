package mongo

import "gopkg.in/mgo.v2/bson"

type Query bson.M

func (q Query) Id(id interface{}) {
	var oid bson.ObjectId

	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			oid = bson.ObjectIdHex(t)
		}
	case bson.ObjectId:
		oid = t
	}

	if oid.Valid() {
		q["_id"] = oid
	}
}

func (q Query) In(field string, value interface{}) {
	q[field] = bson.M{"$in": []interface{}{value}}
}

func (q Query) Push(field string, value interface{}) {
	b := bson.M{}
	b[field] = value

	q["$push"] = b
}

func (q Query) Pull(field string, value interface{}) {
	b := bson.M{}
	b[field] = value

	q["$pull"] = b
}
