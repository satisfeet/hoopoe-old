package store

import (
	"errors"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

// Query is a helper for common mgo queries.
type Query bson.M

var ErrBadParam = errors.New("bad param")

// Sets an equals id condition if id is valid else error is returned.
func (q Query) Id(id interface{}) error {
	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			q["_id"] = bson.ObjectIdHex(t)

			return nil
		}
	case bson.ObjectId:
		if t.Valid() {
			q["_id"] = t

			return nil
		}
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
