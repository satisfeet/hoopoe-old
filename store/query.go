package store

import (
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

// Query is a helper for common mgo queries.
type Query bson.M

// Sets an equals id condition if id is valid else error is returned.
func (q Query) Id(id interface{}) error {
	if oid := ParseId(id); oid.Valid() {
		q["_id"] = oid

		return nil
	}

	return ErrBadParam
}

// Sets an in array condition for id types.
func (q Query) HasId(field string, id interface{}) error {
	if oid := ParseId(id); oid.Valid() {
		q[field] = bson.M{"$in": []bson.ObjectId{oid}}

		return nil
	}

	return ErrBadParam
}

// Fakes a search condition by using or and regex conditions.
// Returns error if query string is invalid which may be ignored for find all.
func (q Query) Search(query string, model interface{}) error {
	if len(query) == 0 {
		return ErrBadParam
	}

	q["$or"] = make([]bson.M, 0)

	for k, _ := range utils.GetStructInfo(model) {
		m := bson.M{}
		m[k] = bson.RegEx{query, "i"}

		q["$or"] = append(q["$or"].([]bson.M), m)
	}

	return nil
}
