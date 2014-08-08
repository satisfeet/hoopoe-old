package store

import "gopkg.in/mgo.v2/bson"

type Update bson.M

func (u Update) PushId(field string, id interface{}) error {
	if oid := ParseId(id); oid.Valid() {
		q := bson.M{}
		q[field] = oid

		u["$push"] = q

		return nil
	}

	return ErrBadParam
}

func (u Update) PullId(field string, id interface{}) error {
	if oid := ParseId(id); oid.Valid() {
		b := bson.M{}
		b[field] = oid

		u["$pull"] = b

		return nil
	}

	return ErrBadParam
}
