package store

import (
	"errors"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
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

type Query bson.M

func (q Query) Id(id bson.ObjectId) {
	q["_id"] = id
}

func (q Query) In(field string, value interface{}) {
	q[field] = bson.M{"$in": []interface{}{value}}
}

func (q Query) Search(query string, model interface{}) {
	if len(query) > 0 {
		q["$or"] = make([]bson.M, 0)

		for k, _ := range utils.GetStructInfo(model) {
			m := bson.M{}
			m[k] = bson.RegEx{query, "i"}

			q["$or"] = append(q["$or"].([]bson.M), m)
		}
	}
}

type Update bson.M

func (u Update) Push(field string, value interface{}) {
	b := bson.M{}
	b[field] = value

	u["$push"] = b
}

func (u Update) Pull(field string, value interface{}) {
	b := bson.M{}
	b[field] = value

	u["$pull"] = b
}
