package mongo

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type Query bson.M

var (
	ErrBadQueryParam = errors.New("bad query id")
	ErrBadQueryValue = errors.New("bad query value")
)

func (query Query) Id(id interface{}) error {
	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			query["_id"] = bson.ObjectIdHex(t)

			return nil
		}
	case bson.ObjectId:
		if t.Valid() {
			query["_id"] = t

			return nil
		}
	}

	return ErrBadQueryParam
}

func (query Query) Or(q Query) error {
	if query["$or"] == nil {
		query["$or"] = make([]Query, 0)
	}
	if or, ok := query["$or"].([]Query); ok {
		query["$or"] = append(or, q)

		return nil
	}

	return ErrBadQueryValue
}

func (query Query) Regex(key, pattern string) error {
	if len(key) == 0 || len(pattern) == 0 {
		return ErrBadQueryParam
	}
	query[key] = bson.RegEx{pattern, "i"}

	return nil
}
